package e

// TODO-xz 后续将ecode与emsg绑定，常用的几个绑定

const (
	// 注册 10000-10100
	ERROR_CODE_REGISTER            = 10000
	ERROR_CODE_REGISTER_PARAMS     = 10001
	ERROR_CODE_AUTH_CONN           = 10002
	ERROR_CODE_AUTH_GENERATE_TOKEN = 10003
)

const (
	ERROR_MSG_REGISTER        = "注册出错"
	ERROR_MSG_REGISTER_PARAMS = "入参错误"
	ERROR_MSG_AUTH_CONN       = "auth服务链接失败"
	ERROR_AUTH_GENERATE_TOKEN = "生成token失败"
)
