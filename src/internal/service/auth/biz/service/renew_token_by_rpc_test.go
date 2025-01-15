package service

import (
	"context"
	auth "github.com/FinnTew/FinnEcommerce/src/internal/service/auth/kitex_gen/auth"
	"testing"
)

func TestRenewTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRenewTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.RenewTokenReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
