package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime/debug"
	"time"
)

type RetryError struct {
	Msg string
}

func (r RetryError) Error() string {
	return r.Msg
}

func NewRetryError(msgParam string) RetryError {
	return RetryError{
		Msg: msgParam,
	}
}

func DoWithRetry(f func() error, times, waitSecond int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("[DoWithRetry] service panic: %v, stack:\r\n%s", r, string(debug.Stack()))
			err = fmt.Errorf("[DoWithRetry] service panic: %v", r)
		}
	}()
	if times <= 1 {
		times = 1
	}
	for i := 1; i <= times; i++ {
		err = f()
		if err == nil {
			return nil
		}
		if err != nil {
			_, ok := err.(RetryError)
			if ok {
				//可重试的错误
				logrus.Infof("[DoWithRetry] do failed with %d times, retry, err: %v", i, err)
			} else {
				logrus.Errorf("[DoWithRetry] do failed with %d times, stop, err: %v", i, err)
				break
			}
			if waitSecond > 0 {
				time.Sleep(time.Duration(waitSecond) * time.Second)
			}
		}
	}
	logrus.Errorf("[DoWithRetry] finally failed with %d times, return, err: %v", times, err)
	return
}
