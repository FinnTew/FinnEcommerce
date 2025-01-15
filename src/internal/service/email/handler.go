package main

import (
	"context"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/email/biz/service"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/email/kitex_gen/email"
)

// EmailServiceImpl implements the last service interface defined in the IDL.
type EmailServiceImpl struct{}

// SendEmail implements the EmailServiceImpl interface.
func (s *EmailServiceImpl) SendEmail(ctx context.Context, req *email.SendEmailRequest) (resp *email.SendEmailResponse, err error) {
	resp, err = service.NewSendEmailService(ctx).Run(req)

	return resp, err
}
