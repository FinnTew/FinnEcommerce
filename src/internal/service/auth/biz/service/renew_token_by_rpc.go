package service

import (
	"context"
	"fmt"
	"github.com/FinnTew/FinnEcommerce/src/internal/pkg/util"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/auth/conf"
	auth "github.com/FinnTew/FinnEcommerce/src/internal/service/auth/kitex_gen/auth"
	"time"
)

type RenewTokenByRPCService struct {
	ctx context.Context
} // NewRenewTokenByRPCService new RenewTokenByRPCService
func NewRenewTokenByRPCService(ctx context.Context) *RenewTokenByRPCService {
	return &RenewTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *RenewTokenByRPCService) Run(req *auth.RenewTokenReq) (resp *auth.RenewTokenResp, err error) {
	token := req.GetToken()

	jwtUtil := util.NewJWTUtil(
		conf.GetConf().Jwt.Secret,
		time.Duration(conf.GetConf().Jwt.Expire)*time.Second,
	)

	refreshToken, err := jwtUtil.RefreshToken(token)
	if err != nil {
		return &auth.RenewTokenResp{
			NewToken: "",
		}, fmt.Errorf("renewTokenByRPCService.Run err: %v", err)
	}

	resp = &auth.RenewTokenResp{
		NewToken: refreshToken,
	}

	return
}
