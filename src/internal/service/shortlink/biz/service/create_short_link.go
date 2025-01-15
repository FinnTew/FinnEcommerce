package service

import (
	"context"
	shortlink "github.com/FinnTew/FinnEcommerce/src/internal/service/shortlink/kitex_gen/shortlink"
)

type CreateShortLinkService struct {
	ctx context.Context
} // NewCreateShortLinkService new CreateShortLinkService
func NewCreateShortLinkService(ctx context.Context) *CreateShortLinkService {
	return &CreateShortLinkService{ctx: ctx}
}

// Run create note info
func (s *CreateShortLinkService) Run(req *shortlink.CreateShortLinkRequest) (resp *shortlink.CreateShortLinkResponse, err error) {
	// Finish your business logic.

	return
}
