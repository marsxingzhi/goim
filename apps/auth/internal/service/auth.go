package service

import (
	"context"
	"fmt"
	"github.com/marsxingzhi/goim/apps/auth/internal/config"
	"github.com/marsxingzhi/goim/apps/auth/internal/repo"
	"github.com/marsxingzhi/goim/domain/model"
	"github.com/marsxingzhi/goim/pkg/common/xzjwt"
	"github.com/marsxingzhi/goim/pkg/common/xzmysql"
	"github.com/marsxingzhi/goim/pkg/common/xzredis"
	"github.com/marsxingzhi/goim/pkg/constant"
	"github.com/marsxingzhi/goim/pkg/e"
	"github.com/marsxingzhi/goim/pkg/proto/pb_auth"
	"github.com/marsxingzhi/goim/pkg/proto/pb_user"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"strconv"
)

// AuthService 服务
type AuthService interface {
	Register(ctx context.Context, req *pb_auth.RegisterReq) (resp *pb_auth.RegisterResp, err error)
	Login(ctx context.Context, req *pb_auth.LoginReq) (resp *pb_auth.LoginResp, err error)
	Logout(ctx context.Context, req *pb_auth.LogoutReq) (resp *pb_auth.LogoutResp, err error)
	RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp, err error)
}

type authService struct {
	conf *config.Config
}

func NewAuthService(conf *config.Config) AuthService {
	return &authService{conf: conf}
}

func (as *authService) Register(ctx context.Context, req *pb_auth.RegisterReq) (resp *pb_auth.RegisterResp, err error) {
	//resp = new(pb_auth.RegisterResp)
	//resp.Msg = "注册成功啦，嘿嘿"
	//resp.Code = 0
	//resp.UserInfo = new(pb_user.UserInfo)
	//
	resp = &pb_auth.RegisterResp{
		UserInfo:     &pb_user.UserInfo{},
		AccessToken:  &pb_auth.Token{},
		RefreshToken: &pb_auth.Token{},
	}

	// 数据库对象
	user := transformUser(req)

	// 入库
	err = xzmysql.Transaction(func(db *gorm.DB) (err error) {
		// 下面交给repo
		authRepo := repo.NewAuthRepo()
		err = authRepo.TxCreate(db, user)
		if err != nil {
			resp.Code = e.ERROR_CODE_REGISTER
			resp.Msg = e.ERROR_MSG_REGISTER
			return
		}
		return
	})

	if err != nil {
		return
	}

	// 生成token
	access, refresh, err := as.generateToken(user.Uid, req.Platform)
	if err != nil {
		resp.Code = e.ERROR_CODE_AUTH_GENERATE_TOKEN
		resp.Msg = e.ERROR_AUTH_GENERATE_TOKEN
		return
	}

	// 返回
	resp.UserInfo = modelUser2pbUser(user)
	resp.AccessToken = access
	resp.RefreshToken = refresh

	fmt.Println("[auth-server] register success...")
	return
}

func (as *authService) Login(ctx context.Context, req *pb_auth.LoginReq) (resp *pb_auth.LoginResp, err error) {
	return
}

func (as *authService) Logout(ctx context.Context, req *pb_auth.LogoutReq) (resp *pb_auth.LogoutResp, err error) {
	return
}

func (as *authService) RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp, err error) {
	return
}

func (as *authService) generateToken(uid int64, platform int32) (access *pb_auth.Token, refresh *pb_auth.Token, err error) {
	var (
		eg               errgroup.Group
		accessSessionId  string
		refreshSessionId string
	)
	access = new(pb_auth.Token)
	refresh = new(pb_auth.Token)

	eg.Go(func() error {
		accessJwtToken, err := xzjwt.GenerateAccessToken(uid, int8(platform), 7*24*3600) // 7天
		if err == nil {
			access.Token = accessJwtToken.Token
			access.Expire = accessJwtToken.Expire
			accessSessionId = accessJwtToken.SessionId
		}
		return err
	})

	eg.Go(func() error {
		refreshJwtToken, err := xzjwt.GenerateRefreshToken(uid, int8(platform), 30*24*3600) // 30天
		if err == nil {
			refresh.Token = refreshJwtToken.Token
			refresh.Expire = refreshJwtToken.Expire
			refreshSessionId = refreshJwtToken.SessionId
		}
		return err
	})

	if err = eg.Wait(); err != nil {
		return nil, nil, fmt.Errorf("failed to generate token: %v", err)
	}

	// sessionId 存储到redis
	key := strconv.FormatInt(uid, 10) + strconv.Itoa(int(platform))
	mp := make(map[string]string)

	k1 := as.conf.Redis.Prefix + constant.REDIS_KEY_USER_ACCESS_TOKEN_SESSION_ID + key
	k2 := as.conf.Redis.Prefix + constant.REDIS_KEY_USER_REFRESH_TOKEN_SESSION_ID + key
	mp[k1] = accessSessionId
	mp[k2] = refreshSessionId
	if err := xzredis.MSet(mp); err != nil {
		// 不要返回
		fmt.Println("[auth] failed to save sessionId to redis: ", err)
	}
	return
}

// TODO-xz  缺少不少值
func transformUser(req *pb_auth.RegisterReq) *model.User {
	var user = new(model.User)
	//user.Uid = uint64(time.Now().Unix()) // TODO-xz 雪花算法  不能再这里就生成，得在入库时
	//user.Aid
	user.Password = req.GetPassword()
	//user.Did
	//user.Status
	user.Nickname = req.GetNickname()
	user.Firstname = req.GetFirstname()
	user.Lastname = req.GetLastname()
	user.Gender = int8(req.GetGender())
	//user.Birth

	user.Email = req.GetEmail()
	user.Mobile = req.GetMobile()
	user.ServerId = 1 // TODO-xz 暂时固定住
	//user.CityId
	//user.AvatarKey

	return user
}

func modelUser2pbUser(u *model.User) *pb_user.UserInfo {
	var pbUser = new(pb_user.UserInfo)
	pbUser.Uid = u.Uid
	pbUser.Status = u.Status
	pbUser.Nickname = u.Nickname
	pbUser.Firstname = u.Firstname
	pbUser.Lastname = u.Lastname
	pbUser.Gender = int32(u.Gender)
	pbUser.BirthTs = u.Birth
	pbUser.Mobile = u.Mobile
	return pbUser
}

func (as *authService) cacheToken(access *pb_auth.Token, refresh *pb_auth.Token, key string) {

}
