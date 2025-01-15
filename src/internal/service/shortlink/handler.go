package main

import (
	"context"
	"github.com/FinnTew/FinnEcommerce/src/internal/service/shortlink/biz/service"
	
)

// ShortLinkServiceImpl implements the last service interface defined in the IDL.
type ShortLinkServiceImpl struct{}

// CreateShortLink implements the ShortLinkServiceImpl interface.
func (s *ShortLinkServiceImpl) CreateShortLink(ctx context.Context, req *shortlink.CreateShortLinkRequest) (resp *shortlink.CreateShortLinkResponse, err error) {
	resp, err = service.NewCreateShortLinkService(ctx).Run(req)

	return resp, err
}

// GetLongLink implements the ShortLinkServiceImpl interface.
func (s *ShortLinkServiceImpl) GetLongLink(ctx context.Context, req *shortlink.GetLongLinkRequest) (resp *shortlink.GetLongLinkResponse, err error) {
	resp, err = service.NewGetLongLinkService(ctx).Run(req)

	return resp, err
}
