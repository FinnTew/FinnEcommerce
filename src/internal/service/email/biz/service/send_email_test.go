package service

import (
	"context"
	email "github.com/FinnTew/FinnEcommerce/src/internal/service/email/kitex_gen/email"
	"testing"
)

func TestSendEmail_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendEmailService(ctx)
	// init req and assert value

	req := &email.SendEmailRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
