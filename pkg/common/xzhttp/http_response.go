package xzhttp

type Result struct {
	Code int32  `json:"code"` // 0：返回成功
	Msg  string `json:"msg"`
}

type Response struct {
	Result
	Data interface{}
}
