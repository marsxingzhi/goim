package xzjwt

import (
	"errors"
	"fmt"
	"github.com/marsxingzhi/goim/pkg/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	JWT_SECRET = "marsxingzhi.goim.2023" // 密钥
	JWT_ISSUER = "marsxingzhi"
	JWT_ACCESS = "jwt_access="
)

type JwtToken struct {
	Token     string
	Expire    int64 // 过期时间
	SessionId string
	Uid       int64
	Platform  int8
}

type XzClaims struct {
	Uid      int64 `json:"uid"`
	Platform int8  `json:"platform"`
	// UserName string `json:"username"`
	// Password string `json:"password"`

	// 由于token是无状态的，那么只要是在有效期内都是有效的，那么就会存在注销登录等场景下token还有效的问题。
	// 因此这里是sessionID保存到redis中，进行二次验证。
	SessionID string `json:"session_id"`
	Access    bool   `json:"access"`
	jwt.StandardClaims
}

func GenerateAccessToken(uid int64, platform int8, duration int) (*JwtToken, error) {
	return generateToken(uid, platform, duration, true)
}

func GenerateRefreshToken(uid int64, platform int8, duration int) (*JwtToken, error) {
	return generateToken(uid, platform, duration, false)
}

func generateToken(uid int64, platform int8, duration int, access bool) (*JwtToken, error) {
	now := time.Now()
	claims := XzClaims{
		Uid:       uid,
		Platform:  platform,
		SessionID: utils.MD5(uuid.NewString()),
		Access:    access,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(time.Second * time.Duration(duration)).Unix(), // 过期时间
			Issuer:    JWT_ISSUER,                                            // 签发人
			IssuedAt:  now.Unix(),                                            // 签发时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(JWT_SECRET))
	if access {
		// 如果是access_token，再加点
		// TODO 这里如果加了，那么在还得使用原始的tokenStr去解析token，然后发现sessionId是refresh的sessionId
		// 因此，还是不要在这里加
		//tokenStr = JWT_ACCESS + tokenStr
	}

	jwtToken := &JwtToken{
		Token:     tokenStr,
		Expire:    claims.ExpiresAt,
		SessionId: claims.SessionID,
		Uid:       uid,
		Platform:  platform,
	}

	return jwtToken, err
}

func ParseToken(tokenStr string) (*XzClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &XzClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		fmt.Println("[xzjwt] failed to parsewithclaims, err: ", err.Error())
		return nil, err
	}

	// 验证token
	if claims, ok := token.Claims.(*XzClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("this is valid token")
}

func ParseFromAuthorization(ctx *gin.Context) (*XzClaims, error) {
	tokenStr := ctx.GetHeader("Authorization")
	return ParseToken(tokenStr)
}
