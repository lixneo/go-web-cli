package controller

type MyCode int64

const (
	CodeSuccess           MyCode = 1
	CodeServerBusy        MyCode = 500
	CodeInvalidParams     MyCode = 1001
	CodeUserExist         MyCode = 1002
	CodeUserNotExist      MyCode = 1003
	CodeInvalidPassword   MyCode = 1004
	CodeInvalidToken      MyCode = 1005
	CodeInvalidAuthFormat MyCode = 1006
	CodeNotLogin          MyCode = 1007
)

var msgFlags = map[MyCode]string{
	CodeSuccess:           "success",
	CodeServerBusy:        "服务繁忙",
	CodeInvalidParams:     "请求参数错误",
	CodeUserExist:         "用户名重复",
	CodeUserNotExist:      "用户不存在",
	CodeInvalidPassword:   "用户名或密码错误",
	CodeInvalidToken:      "无效的Token",
	CodeInvalidAuthFormat: "认证格式有误",
	CodeNotLogin:          "未登录",
}

func (c MyCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
