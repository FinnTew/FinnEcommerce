package service

import (
	"context"
	"fmt"
	authk "github.com/FinnTew/FinnEcommerce/src/internal/client/kitex_gen/auth"
	"github.com/FinnTew/FinnEcommerce/src/internal/pkg/util"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/client"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/model"
	user "github.com/FinnTew/FinnEcommerce/src/internal/service/user/kitex_gen/user"
)

type UpdateService struct {
	ctx context.Context
} // NewUpdateService new UpdateService
func NewUpdateService(ctx context.Context) *UpdateService {
	return &UpdateService{ctx: ctx}
}

// Run create note info
func (s *UpdateService) Run(req *user.UpdateReq) (resp *user.UpdateResp, err error) {
	userID := req.GetUserId()
	email := req.GetEmail()
	password := req.GetPassword()
	token := req.GetAccessToken()
	if email == "" {
		return nil, fmt.Errorf("updateService.Run err: email is empty")
	}
	if password == "" {
		return nil, fmt.Errorf("updateService.Run err: password is empty")
	}

	verifyTokenResp, err := client.AuthClient.VerifyTokenByRPC(s.ctx, &authk.VerifyTokenReq{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	if !verifyTokenResp.GetValid() {
		return nil, fmt.Errorf("updateService.Run err: invalid token")
	}

	u, err := model.GetUserByID(mysql.DB, s.ctx, uint32(userID))
	if err != nil {
		return nil, fmt.Errorf("updateService.Run err: %v", err)
	}

	if u.Email != email {
		u.Email = email
	}
	if err = util.NewHashPwdUtil().ValidatePwd(u.Password, password); err != nil {
		hashedPassword, err := util.NewHashPwdUtil().GenerateHash(password)
		if err != nil {
			return nil, fmt.Errorf("updateService.Run err: %v", err)
		}
		u.Password = hashedPassword
	}

	err = model.UpdateUser(mysql.DB, s.ctx, u)
	if err != nil {
		return nil, fmt.Errorf("updateService.Run err: %v", err)
	}

	resp = &user.UpdateResp{
		Success: true,
	}

	return
}
