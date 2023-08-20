package conf

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Config AppConfig
)

type AppConfig struct {
	YamlConfig
}

type YamlConfig struct {
	Port     int
	LogLevel string
	DSN      string
	Env      string
	Cookie   struct {
		Domain string
	}
	Redis struct {
		Addr string
		Pwd  string
		DB   int
	}
	OpenAI struct {
		AuthKey            string
		ChatCompletionsUrl string
		Model              string
		Temperature        float64
		MaxTokens          int
	}
	HttpProxy struct {
		Url string
	}
	Messenger struct {
		AccessToken string
		VerifyToken string
	}
}

func LoadConf() {
	parseFlags()
	parseConf()
}

func parseFlags() {
	e := flag.String("env", "", "env include prod|dev")
	flag.Parse()
	if *e != "" {
		logrus.Infof("set env by flag. env:%s", *e)
		_ = os.Setenv("env", *e)
	}
}

func parseConf() {
	v := viper.New()
	v.SetEnvPrefix("YC")

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")
	err := v.ReadInConfig()
	if err != nil {
		logrus.Panicf("Fatal error read config file: %s \n", err)
	}
	os.Getenv("")
	mode := os.Getenv("env")
	if mode == "" {
		mode = "dev"
	}
	vv := v.Sub(mode)
	yamlConfig := &Config.YamlConfig
	if err = vv.Unmarshal(yamlConfig); err != nil {
		logrus.Panicf("Fatal error unmarshal config file: %s \n", err)
	}
	yamlConfig.Env = mode
	logrus.Infof("load config. %v", yamlConfig)
}
