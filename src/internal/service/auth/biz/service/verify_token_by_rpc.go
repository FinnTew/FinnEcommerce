package service

import (
	"context"
	"fmt"
	"github.com/FinnTew/FinnEcommerce/src/internal/pkg/util"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/auth/conf"
	auth "github.com/FinnTew/FinnEcommerce/src/internal/service/auth/kitex_gen/auth"
	"time"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	token := req.GetToken()

	jwtUtil := util.NewJWTUtil(
		conf.GetConf().Jwt.Secret,
		time.Duration(conf.GetConf().Jwt.Expire)*time.Second,
	)
	claims, err := jwtUtil.ParseToken(token)
	if err != nil {
		return &auth.VerifyResp{
			UserId: "0",
			Valid:  false,
		}, fmt.Errorf("verifyTokenByRPCService.Run err: %v", err)
	}

	resp = &auth.VerifyResp{
		UserId: claims.UserID,
		Valid:  true,
	}

	return

}
