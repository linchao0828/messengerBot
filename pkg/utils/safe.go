package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"runtime/debug"
	"time"
)

func ProtectFn(c context.Context, fn func(), fallback func()) {
	defer func() {
		if r := recover(); r != nil {
			logrus.WithContext(c).Errorf("handle fn panic! %s", string(debug.Stack()))
			time.Sleep(1 * time.Second)
			fallback()
		}
	}()
	fn()
}
