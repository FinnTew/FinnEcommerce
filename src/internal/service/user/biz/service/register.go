package service

import (
	"context"
	"fmt"
	"github.com/FinnTew/FinnEcommerce/src/internal/pkg/util"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/model"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	email := req.GetEmail()
	password := req.GetPassword()
	comfirmPassword := req.GetConfirmPassword()

	if password != comfirmPassword {
		return nil, fmt.Errorf("registerService.Run err: password and comfirmPassword not match")
	}

	hashedPassword, err := util.NewHashPwdUtil().GenerateHash(password)
	if err != nil {
		return nil, fmt.Errorf("RegisterService.Run err: %v", err)
	}

	// TODO: add permissions for common user
	u := &model.User{
		Email:           email,
		Password:        hashedPassword,
		IsEmailVerified: false,
		Role:            "user",
		Permissions:     "{}",
	}
	err = model.CreateUser(mysql.DB, s.ctx, u)
	if err != nil {
		return nil, fmt.Errorf("RegisterService.Run err: %v", err)
	}

	// TODO: verify the email for registered user,
	//       need to implement the email service first

	resp = &user.RegisterResp{
		UserId: int32(u.ID),
	}

	return
}
