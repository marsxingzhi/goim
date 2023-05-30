package dto_auth

// TODO-xz 参数校验，可以通过标签
// 参考 https://blog.csdn.net/weixin_41922289/article/details/121432884
type RegisterReq struct {
	Platform  int32  `json:"platform,omitempty" form:"platform"`                    // 注册平台 0：unknown；1：ios；2：android；3：web
	Nickname  string `json:"nickname,omitempty" form:"nickname" binding:"required"` // 昵称
	Password  string `json:"password,omitempty" form:"password" binding:"required"` // 密码
	Firstname string `json:"firstname,omitempty" form:"firstname"`
	Lastname  string `json:"lastname,omitempty" form:"lastname"`
	Gender    int32  `json:"gender,omitempty" form:"gender"` // 1：男;；2：女
	Email     string `json:"email,omitempty" form:"email"`   // 邮箱
	Mobile    string `json:"mobile,omitempty" form:"mobile"` // 手机号
}
