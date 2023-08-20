package yerrors

import (
	"fmt"
	"github/linchao0828/messengerBot/pkg/consts"
	"net/http"
)

var (
	BadRequest = BizError{
		Status: http.StatusBadRequest,
		Code:   consts.BadRequest,
		Msg:    "参数错误",
	}
	StatusUnauthorized = BizError{
		Status: http.StatusUnauthorized,
		Code:   consts.NoAuthorized,
		Msg:    "请登录后再试",
	}
	RecordNotFound = BizError{
		Status: http.StatusNotFound,
		Code:   consts.NotFound,
		Msg:    "数据不存在，请检查",
	}
	ServerError = BizError{
		Status: http.StatusInternalServerError,
		Code:   consts.Fail,
		Msg:    "系统错误，请稍后重试",
	}
	RecordAlreadyExist = BizError{
		Status: http.StatusConflict,
		Code:   consts.AlreadyExist,
		Msg:    "数据已存在，请勿重复提交",
	}
	MobileCodeVerifyFail = BizError{
		Status: http.StatusBadRequest,
		Code:   consts.MobileCodeVerifyFail,
		Msg:    "手机验证码验证失败",
	}
	UserNotFound = BizError{
		Status: http.StatusNotFound,
		Code:   consts.UserNotFound,
		Msg:    "用户不存在",
	}
	ContentIllegal = BizError{
		Status: http.StatusBadRequest,
		Code:   consts.ContentIllegal,
		Msg:    "抱歉，内容包含敏感词",
	}
)

type BizError struct {
	Status   int
	Code     int
	Msg      string
	MsgParam []interface{}
}

func (b BizError) Error() string {
	return fmt.Sprintf(b.Msg, b.MsgParam...)
}

func FnInternalServerError(msg string) BizError {
	return BizError{
		Status: http.StatusInternalServerError,
		Code:   consts.Fail,
		Msg:    msg,
	}
}

func FnErrorWithMsgParam(err BizError, msgParam ...interface{}) BizError {
	err.MsgParam = msgParam
	return err
}
