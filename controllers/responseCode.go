package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNoExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken

	CodeInsertFailed

	CodePostInvalid

	CodeDataIsNull
	CodeRateLimit
)

var codeMsg = map[ResCode]string{
	CodeSuccess:         "Success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNoExist:     "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务器繁忙",

	CodeNeedLogin:    "请先登录",
	CodeInvalidToken: "无效的Token",

	CodeInsertFailed: "写入数据失败",

	CodePostInvalid: "当前PostID错误",

	CodeDataIsNull: "查询结果为空",
	CodeRateLimit:  "当前访问人数过多，请稍后再尝试噢...",
}

func (code ResCode) Msg() (msg string) {
	msg, ok := codeMsg[code]
	if !ok {
		msg = codeMsg[CodeServerBusy]
	}
	return msg
}
