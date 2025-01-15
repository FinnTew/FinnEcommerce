package service

import (
	"context"
	user "github.com/FinnTew/FinnEcommerce/src/internal/service/user/kitex_gen/user"
	"testing"
)

func TestDelete_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteService(ctx)
	// init req and assert value

	req := &user.DeleteReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
