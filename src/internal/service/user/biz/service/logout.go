package service

import (
	"context"
	"fmt"
	authk "github.com/FinnTew/FinnEcommerce/src/internal/client/kitex_gen/auth"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/client"
	user "github.com/FinnTew/FinnEcommerce/src/internal/service/user/kitex_gen/user"
)

type LogoutService struct {
	ctx context.Context
} // NewLogoutService new LogoutService
func NewLogoutService(ctx context.Context) *LogoutService {
	return &LogoutService{ctx: ctx}
}

// Run create note info
func (s *LogoutService) Run(req *user.LogoutReq) (resp *user.LogoutResp, err error) {
	token := req.GetAccessToken()

	verifyTokenResp, err := client.AuthClient.VerifyTokenByRPC(s.ctx, &authk.VerifyTokenReq{
		Token: token,
	})
	if err != nil {
		return nil, fmt.Errorf("logoutService.Run err: %v", err)
	}
	if !verifyTokenResp.GetValid() {
		return nil, fmt.Errorf("logoutService.Run err: invalid token")
	}

	resp = &user.LogoutResp{
		Success: true,
	}

	return
}
