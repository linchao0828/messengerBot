package openai

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github/linchao0828/messengerBot/domain"
	"github/linchao0828/messengerBot/pkg/logger"
	"github/linchao0828/messengerBot/pkg/utils"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	OpenAIError_ServerError = "server_error"
	OpenAIError_LengthError = "context_length_exceeded"
)

var (
	once sync.Once
	Cli  *client
)

type client struct {
	cli                                *http.Client
	authKey, chatCompletionsUrl, model string
	temperature                        float64
	maxTokens                          int
}

// ChatGPTResponseBody 响应体
type ChatGPTResponseBody struct {
	ID      string                `json:"id"`
	Object  string                `json:"object"`
	Created int                   `json:"created"`
	Model   string                `json:"model"`
	Choices []Choices             `json:"choices"`
	Usage   Usage                 `json:"usage"`
	Error   *ChatGPTResponseError `json:"error"`
}

type ChatGPTResponseError struct {
	Code    string `json:"code"`
	Param   string `json:"param"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Choices struct {
	Message      domain.ChatMessage `json:"message"`
	Delta        domain.ChatMessage `json:"delta"`
	FinishReason string             `json:"finish_reason"`
	Index        int                `json:"index"`
}

// ChatGPTRequestBody 请求体
type ChatGPTRequestBody struct {
	Model            string                `json:"model"`
	Prompt           *string               `json:"prompt,omitempty"`
	Messages         []*domain.ChatMessage `json:"messages,omitempty"`
	MaxTokens        *int                  `json:"max_tokens,omitempty"`
	Temperature      *float64              `json:"temperature,omitempty"`
	TopP             *int                  `json:"top_p,omitempty"`
	FrequencyPenalty *int                  `json:"frequency_penalty,omitempty"`
	PresencePenalty  *int                  `json:"presence_penalty,omitempty"`
	Stop             []string              `json:"stop,omitempty"`
	Stream           bool                  `json:"stream,omitempty"`
}

func Init(authKey, chatCompletionsUrl, model string, temperature float64, maxTokens int, proxyUrl string) {
	once.Do(func() {
		proxy := http.ProxyFromEnvironment
		if proxyUrl != "" {
			proxy = func(req *http.Request) (*url.URL, error) {
				return url.Parse(proxyUrl)
			}
		}
		Cli = &client{
			authKey:            authKey,
			chatCompletionsUrl: chatCompletionsUrl,
			model:              model,
			temperature:        temperature,
			maxTokens:          maxTokens,
			cli: &http.Client{
				Timeout: 30 * time.Second,
				Transport: &http.Transport{
					Proxy: proxy,
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			},
		}
	})
}

// Chat
func (c *client) Chat(ctx *gin.Context, chatMessageList []*domain.ChatMessage) (*domain.ChatMessage, int, error) {
	if len(chatMessageList) == 0 {
		return nil, 0, errors.New("chatMessageList is empty")
	}

	logger.WithContext(ctx).WithField("chatMessageList", chatMessageList).Info("openai Chat request message")
	requestBody := ChatGPTRequestBody{
		Model:    c.model,
		Messages: chatMessageList,
		//MaxTokens:   utils.IntToPtr(c.maxTokens),
		Temperature: utils.Float64ToPtr(c.temperature),
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("openai Chat error")
		return nil, 0, err
	}
	req, err := http.NewRequest("POST", c.chatCompletionsUrl, bytes.NewBuffer(requestData))
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("openai Chat error")
		return nil, 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authKey))
	response, err := c.cli.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}
	logger.WithContext(ctx).WithField("body", string(body)).Info("openai Chat response body")

	gptResponseBody := &ChatGPTResponseBody{}
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("openai Chat error")
		return nil, 0, err
	}

	if gptResponseBody.Error != nil {
		err = errors.New(gptResponseBody.Error.Message)
		logger.WithContext(ctx).WithError(err).Error("openai Chat response error")
		//openai的服务错误，返回重试，且不删减token
		if gptResponseBody.Error.Type == "server_error" {
			return nil, 0, utils.NewRetryError(OpenAIError_ServerError)
		}
		//token超了，返回重试
		if gptResponseBody.Error.Code == "context_length_exceeded" {
			return nil, 0, utils.NewRetryError(OpenAIError_LengthError)
		}
		return nil, 0, err
	}

	var reply string
	if len(gptResponseBody.Choices) > 0 {
		//token超了，返回重试
		if gptResponseBody.Choices[0].FinishReason == "length" {
			return nil, 0, utils.NewRetryError(OpenAIError_LengthError)
		}
		reply = gptResponseBody.Choices[0].Message.Content
	} else {
		return nil, 0, errors.New("reply is empty")
	}
	logger.WithContext(ctx).Infof("openai Chat response prompt_tokens use %d, text: %s", gptResponseBody.Usage.PromptTokens, reply)
	return &domain.ChatMessage{
		Role:    domain.ChatMessageRole_Assistant,
		Content: reply,
	}, gptResponseBody.Usage.TotalTokens, nil
}

// PackChatMessage 打包聊天消息
func (c *client) PackChatMessage(ctx *gin.Context, prompt, chatContent string, historyMessageList []*domain.ChatMessage) []*domain.ChatMessage {
	var chatMessageList []*domain.ChatMessage
	if prompt != "" {
		chatMessageList = append(chatMessageList, &domain.ChatMessage{
			Role:    domain.ChatMessageRole_System,
			Content: prompt,
		})
	}

	if len(historyMessageList) > 0 {
		chatMessageList = append(chatMessageList, historyMessageList...)
	}

	if chatContent != "" {
		chatMessageList = append(chatMessageList, &domain.ChatMessage{
			Role:    domain.ChatMessageRole_User,
			Content: chatContent,
		})
	}

	return chatMessageList
}

// ChatStream
func (c *client) ChatStream(ctx *gin.Context, chatMessageList []*domain.ChatMessage, chatMessageChannel chan *domain.ChatMessage) error {
	if len(chatMessageList) == 0 {
		return errors.New("chatMessageList is empty")
	}

	logger.WithContext(ctx).WithField("chatMessageList", chatMessageList).Info("openai ChatStream request message")
	requestBody := ChatGPTRequestBody{
		Model:    c.model,
		Messages: chatMessageList,
		//MaxTokens:   utils.IntToPtr(c.maxTokens),
		Temperature: utils.Float64ToPtr(c.temperature),
		Stream:      true,
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("openai ChatStream error")
		return err
	}
	req, err := http.NewRequest("POST", c.chatCompletionsUrl, bytes.NewBuffer(requestData))
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("openai ChatStream error")
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authKey))
	req.Header.Set("Accept", "text/event-stream")
	response, err := c.cli.Do(req)
	if err != nil {
		return err
	}

	go func() {
		scanner := bufio.NewScanner(response.Body)
		for scanner.Scan() {
			body := scanner.Text()
			if body == "" {
				continue
			}
			if body == "data: [DONE]" {
				chatMessageChannel <- nil
				//response.Body.Close()
				return
			}

			gptResponseBody := &ChatGPTResponseBody{}
			jsonData := strings.TrimPrefix(body, "data: ")
			err = json.Unmarshal([]byte(jsonData), gptResponseBody)
			if err != nil {
				logger.WithContext(ctx).WithError(err).Error("openai Chat error")
				return
			}

			if len(gptResponseBody.Choices) > 0 && gptResponseBody.Choices[0].Delta.Content != "" {
				chatMessageChannel <- &domain.ChatMessage{
					Role:    domain.ChatMessageRole_Assistant,
					Content: gptResponseBody.Choices[0].Delta.Content,
				}
			}
		}
	}()

	return nil
}
