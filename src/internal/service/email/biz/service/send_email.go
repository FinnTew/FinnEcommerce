package service

import (
	"context"
	email "github.com/FinnTew/FinnEcommerce/src/internal/service/email/kitex_gen/email"
)

type SendEmailService struct {
	ctx context.Context
} // NewSendEmailService new SendEmailService
func NewSendEmailService(ctx context.Context) *SendEmailService {
	return &SendEmailService{ctx: ctx}
}

// Run create note info
func (s *SendEmailService) Run(req *email.SendEmailRequest) (resp *email.SendEmailResponse, err error) {
	// Finish your business logic.

	return
}
