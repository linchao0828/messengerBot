package api

import (
	"github/linchao0828/messengerBot/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github/linchao0828/messengerBot/pkg/consts"
	"github/linchao0828/messengerBot/pkg/yerrors"
)

type BaseResp struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	LogId string      `json:"logId"`
	Data  interface{} `json:"data"`
}

type IdData struct {
	Id interface{} `json:"id"`
}

type PageResp struct {
	Count int64       `json:"count"`
	Items interface{} `json:"items"`
}

type CursorResp struct {
	HasMore bool        `json:"has_more"`
	Items   interface{} `json:"items"`
}

func OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, BaseResp{
		Code:  consts.OK,
		Msg:   "ok",
		LogId: ctx.GetString(consts.LogId),
		Data:  data,
	})
	ctx.Abort()
}

func Fail(ctx *gin.Context, err error) {
	logger.WithContext(ctx).WithError(err).Error("http server error")
	var e yerrors.BizError
	if be, ok := err.(yerrors.BizError); ok {
		e = be
	} else {
		e = yerrors.ServerError
	}

	ctx.JSON(e.Status, BaseResp{
		Code:  e.Code,
		Msg:   e.Error(),
		LogId: ctx.GetString(consts.LogId),
	})
	ctx.Abort()
}

func Page(ctx *gin.Context, items interface{}, count int64) {
	ctx.JSON(http.StatusOK, BaseResp{
		Code:  consts.OK,
		Msg:   "ok",
		LogId: ctx.GetString(consts.LogId),
		Data: PageResp{
			Items: items,
			Count: count,
		},
	})
	ctx.Abort()
}

func Cursor(ctx *gin.Context, items interface{}, hasMore bool) {
	ctx.JSON(http.StatusOK, BaseResp{
		Code:  consts.OK,
		Msg:   "ok",
		LogId: ctx.GetString(consts.LogId),
		Data: CursorResp{
			Items:   items,
			HasMore: hasMore,
		},
	})
	ctx.Abort()
}

func Write(ctx *gin.Context, status int, code int, msg string) {
	ctx.JSON(status, BaseResp{
		Code:  code,
		Msg:   msg,
		LogId: ctx.GetString(consts.LogId),
	})
	ctx.Abort()
}

func OKWithString(ctx *gin.Context, str string) {
	ctx.String(http.StatusOK, str)
	ctx.Abort()
}

func OKWithJson(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
	ctx.Abort()
}
