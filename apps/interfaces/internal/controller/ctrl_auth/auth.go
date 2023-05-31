package ctrl_auth

import "github.com/marsxingzhi/goim/apps/interfaces/internal/service/auth"

//type AuthCtrl interface {
//}
//

type AuthCtrl struct {
	authService auth.AuthService
}

// 这里没必要使用接口

func NewAuthCtrl(ac auth.AuthService) *AuthCtrl {
	return &AuthCtrl{authService: ac}
}
