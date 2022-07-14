package errorcode

var (
    UserLoginFailedCode = StandardError{100001, "用户名或密码错误"}
    TokenExpired        = StandardError{100002, "Token已过期"}
    TokenNotValidYet    = StandardError{100003, "Token尚未激活"}
    TokenMalformed      = StandardError{100004, "非登陆失效Token类型"}
    TokenInvalid        = StandardError{100005, "无效的Token"}
    TokenReplaced       = StandardError{100006, "登录已失效，请重新登录"}
    TokenCreatedFail    = StandardError{100007, "Token生成失败"}
    TokenSomethingError = StandardError{100008, "Token未知错误"}

    UserAddressCreateFailed     = StandardError{101001, "收货地址保存失败"}
    UserAddressNotFound         = StandardError{101002, "收货地址不存在"}
    UserAddressDeleteFailed     = StandardError{101003, "收货地址删除失败"}
    UserAddressDefaultFailed    = StandardError{101004, "设置默认收货地址失败"}
    UserAddressModifyFailed     = StandardError{101005, "收货地址修改失败"}
    UserAddressNotDefaultFailed = StandardError{101006, "收货地址设置非默认失败"}
    UserAddressResolveFailed    = StandardError{101007, "地址解析失败"}

    UserCreatedFail         = StandardError{102001, "用户创建失败"}
    UserLoginTimeUpdateFail = StandardError{102002, "用户最近登录时间更新失败"}

    UserDeviceSetFailed = StandardError{103001, "用户设备信息设置失败"}
)
