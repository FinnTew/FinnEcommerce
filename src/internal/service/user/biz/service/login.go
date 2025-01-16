package service

import (
	"context"
	"fmt"
	authk "github.com/FinnTew/FinnEcommerce/src/internal/client/kitex_gen/auth"
	"github.com/FinnTew/FinnEcommerce/src/internal/pkg/util"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/client"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/dal/redis"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/model"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/kitex_gen/user"
	"time"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	email := req.GetEmail()
	password := req.GetPassword()

	u, err := model.GetUserByEmail(mysql.DB, s.ctx, email)
	if err != nil {
		return nil, fmt.Errorf("loginService.Run err: %v", err)
	}
	if !u.IsEmailVerified {
		return nil, fmt.Errorf("loginService.Run err: email not verified")
	}

	hashPwdUtil := util.NewHashPwdUtil()
	err = hashPwdUtil.ValidatePwd(u.Password, password)
	if err != nil {
		return nil, fmt.Errorf("loginService.Run err: %v", err)
	}

	reliverTokenResp, err := client.AuthClient.DeliverTokenByRPC(s.ctx, &authk.DeliverTokenReq{
		UserId: int32(u.ID),
	})
	if err != nil {
		return nil, err
	}

	token := reliverTokenResp.GetAccessToken()
	expiresIn := reliverTokenResp.GetExpiresIn()

	key := util.NewRedisKey().GetUserTokenKey(u.ID)
	err = redis.RedisClient.Set(s.ctx, key, token, time.Duration(expiresIn)).Err()
	if err != nil {
		return nil, fmt.Errorf("loginService.Run err: %v", err)
	}

	resp = &user.LoginResp{
		UserId:      int32(u.ID),
		AccessToken: token,
		ExpiresIn:   expiresIn,
	}

	return
}
