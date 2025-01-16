package service

import (
	"context"
	"fmt"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/dal/mysql"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/user/biz/model"
	user "github.com/FinnTew/FinnEcommerce/src/internal/service/user/kitex_gen/user"
)

type DeleteService struct {
	ctx context.Context
} // NewDeleteService new DeleteService
func NewDeleteService(ctx context.Context) *DeleteService {
	return &DeleteService{ctx: ctx}
}

// Run create note info
func (s *DeleteService) Run(req *user.DeleteReq) (resp *user.DeleteResp, err error) {
	userID := req.GetUserId()

	_, err = model.GetUserByID(mysql.DB, s.ctx, uint32(userID))
	if err != nil {
		return &user.DeleteResp{
			Success: false,
		}, fmt.Errorf("deleteService.Run err: %v", err)
	}

	err = model.DeleteUserByID(mysql.DB, s.ctx, uint32(userID))
	if err != nil {
		return &user.DeleteResp{
			Success: false,
		}, fmt.Errorf("deleteService.Run err: %v", err)
	}

	resp = &user.DeleteResp{
		Success: true,
	}

	return
}
