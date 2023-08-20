package logger

import (
	"context"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github/linchao0828/messengerBot/pkg/consts"
	"os"
)

var keys = map[string]bool{
	consts.CurrentUserId: true,
}

func Init(level string) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	l, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Panicf("parse log level panic! err:%s", err)
	}
	//path := "./logs/app.log"
	//f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(l)
}

func WithContext(ctx *gin.Context) *logrus.Entry {
	logId := ctx.GetString(consts.LogId)
	if logId == "" {
		logId = uuid.NewV4().String()
		ctx.Set(consts.LogId, logId)
	}
	logEntry := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"log_id": logId,
	})
	for k := range keys {
		v, ok := ctx.Get(k)
		if ok {
			logEntry = logEntry.WithField(k, v)
		}
	}
	return logEntry
}

func WithOriContext(ctx context.Context) *logrus.Entry {
	logId := cast.ToString(ctx.Value(consts.LogId))
	if logId == "" {
		logId = uuid.NewV4().String()
		ctx = context.WithValue(ctx, consts.LogId, logId)
	}
	logEntry := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"log_id": logId,
	})
	for k := range keys {
		logEntry = logEntry.WithField(k, ctx.Value(k))
	}
	return logEntry
}
