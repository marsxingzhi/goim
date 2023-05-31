package dto_auth

import (
	"github.com/marsxingzhi/goim/pkg/proto/pb_auth"
	"github.com/marsxingzhi/goim/pkg/proto/pb_user"
)

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

// 使用dto的好处：这里可以组装数据返回给端上
// 如果直接将pb_user.UserInfo作为data返回，不利于后续扩展；而且如果pb_user.UserInfo字段很多，只想返回部分user字段的话，也只需要修改RegisterResp即可
type RegisterResp struct {
	UserInfo     *pb_user.UserInfo `json:"user_info,omitempty"`
	AccessToken  *pb_auth.Token    `json:"access_token"`
	RefreshToken *pb_auth.Token    `json:"refresh_token"`
}
