package errorcode

import "strings"

type StandardError struct {
	ErrorCode int
	ErrorMsg  string
}

var (
	SuccessError = StandardError{0, "成功"}

	RecordNotFound        = StandardError{8001, "记录未找到"}
	RecordStatusIncorrect = StandardError{8002, "记录状态不正确"}
	RecordUpdateFailed    = StandardError{8003, "记录更新失败"}

	ValidateError                 = StandardError{9000, "请填写信息后在提交"}
	ServiceError                  = StandardError{9001, "服务内部错误"}
	HandlerError                  = StandardError{9002, "处理异常"}
	SystemDatabaseConnectionError = StandardError{9003, "系统数据链接出错"}
	SystemCacheConnectionError    = StandardError{9004, "系统缓存链接出错"}
	TokenInBlackListError         = StandardError{9005, "token 已被加入黑名单"}
	PermissionError               = StandardError{9007, "没有权限访问"}
	EmptyError                    = StandardError{9008, "数据空"}
	HashBeenError                 = StandardError{9009, "数据已操作"}
	RequestSignValidateError      = StandardError{9010, "请求签名参数验证不正确"}
	RequestLoginError             = StandardError{9011, "请重新登陆"}
	VipExpiredError               = StandardError{9012, "VIP已经过期"}
)

func (s StandardError) SetErrorMsg(msg string) StandardError {
	if strings.Contains(msg, "Duplicate") {
		return HandlerError
	} else {
		s.ErrorMsg = msg
	}

	return s
}
