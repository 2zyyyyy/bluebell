package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNoExist
	CodeInvalidPassword
	CodeServerBusy
)

var codeMsg = map[ResCode]string{
	CodeSuccess:         "Success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNoExist:     "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务器繁忙",
}

func (code ResCode) Msg() (msg string) {
	msg, ok := codeMsg[code]
	if !ok {
		msg = codeMsg[CodeServerBusy]
	}
	return msg
}
