package service

import (
	"context"
	"fmt"
	authk "github.com/FinnTew/FinnEcommerce/src/internal/client/kitex_gen/auth"
	"github.com/FinnTew/FinnEcommerce/src/internal/client/rpc/auth"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/auth/conf"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/model"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/kitex_gen/user"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

type GetUserInfoService struct {
	ctx context.Context
} // NewGetUserInfoService new GetUserInfoService
func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

// Run create note info
func (s *GetUserInfoService) Run(req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	userID := req.GetUserId()
	token := req.GetAccessToken()

	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	authClient, err := auth.NewRPCClient("auth", client.WithResolver(r))
	if err != nil {
		return nil, fmt.Errorf("getUserInfoService.Run err: %v", err)
	}

	verifyTokenResp, err := authClient.VerifyTokenByRPC(s.ctx, &authk.VerifyTokenReq{
		Token: token,
	})
	if err != nil {
		return nil, fmt.Errorf("getUserInfoService.Run err: %v", err)
	}
	if verifyTokenResp.GetValid() == false {
		return nil, fmt.Errorf("getUserInfoService.Run err: %v", "invalid token")
	}

	u, err := model.GetUserByID(mysql.DB, s.ctx, uint32(userID))
	if err != nil {
		return nil, fmt.Errorf("getUserInfoService.Run err: %v", err)
	}

	resp = &user.GetUserInfoResp{
		UserId:    int32(u.ID),
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return
}
