package service

import (
	"context"
	shortlink "github.com/FinnTew/FinnEcommerce/src/internal/service/shortlink/kitex_gen/shortlink"
)

type GetLongLinkService struct {
	ctx context.Context
} // NewGetLongLinkService new GetLongLinkService
func NewGetLongLinkService(ctx context.Context) *GetLongLinkService {
	return &GetLongLinkService{ctx: ctx}
}

// Run create note info
func (s *GetLongLinkService) Run(req *shortlink.GetLongLinkRequest) (resp *shortlink.GetLongLinkResponse, err error) {
	// Finish your business logic.

	return
}
