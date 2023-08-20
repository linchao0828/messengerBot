package cron

import (
	"context"
	"github.com/sirupsen/logrus"
	"runtime/debug"
	"time"
)

type Cron struct {
}

func (c Cron) Start() {
	go c.none()
}

func (c Cron) none() {
}

func protectFn(c context.Context, fn func()) {
	defer func() {
		if r := recover(); r != nil {
			logrus.WithContext(c).Errorf("handle fn panic! %s", string(debug.Stack()))
			time.Sleep(1 * time.Second)
		}
	}()
	fn()
}
