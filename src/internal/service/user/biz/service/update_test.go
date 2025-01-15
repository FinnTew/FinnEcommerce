package service

import (
	"context"
	user "github.com/FinnTew/FinnEcommerce/src/internal/service/user/kitex_gen/user"
	"testing"
)

func TestUpdate_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateService(ctx)
	// init req and assert value

	req := &user.UpdateReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
