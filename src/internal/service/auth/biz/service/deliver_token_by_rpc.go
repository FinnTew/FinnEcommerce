package service

import (
	"context"
	"fmt"
	"github.com/FinnTew/FinnEcommerce/src/internal/pkg/util"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/auth/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/auth/biz/model"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/auth/conf"
	auth "github.com/FinnTew/FinnEcommerce/src/internal/service/auth/kitex_gen/auth"
	"strconv"
	"time"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	userID := req.GetUserId()

	user, err := model.GetUserByID(mysql.DB, s.ctx, uint32(userID))
	if err != nil {
		return nil, fmt.Errorf("deliverTokenByRPCService.Run err: %v", err)
	}
	role := user.Role

	jwtUtil := util.NewJWTUtil(
		conf.GetConf().Jwt.Secret,
		time.Duration(conf.GetConf().Jwt.Expire)*time.Second,
	)
	token, err := jwtUtil.GenerateToken(strconv.Itoa(int(userID)), role)
	if err != nil {
		return nil, fmt.Errorf("deliverTokenByRPCService.Run err: %v", err)
	}

	resp = &auth.DeliveryResp{
		AccessToken: token,
		ExpiresIn:   int32(conf.GetConf().Jwt.Expire),
	}

	return
}
