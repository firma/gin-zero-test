package errno

var (
	OK                     = NewError(0, "操作成功")
	ParamValidationErr     = NewError(100, "参数验证不正确")
	RecordNotFound         = NewError(101, "记录未找到")
	RecordNotFoundWithMore = NewError(101, "记录未找到，失败原因 %s")
	MethodNotAllowed       = NewError(102, "请求方法不允许")
	Forbidden              = NewError(111, "无权操作")
	TimeOut                = NewError(120, "请求超时")
	NetworkErr             = NewError(140, "请求超时")
	SysErr                 = NewError(145, "系统错误")
)
