package shortlink

import (
	"context"
	shortlink "github.com/FinnTew/FinnEcommerce/src/internal/client/kitex_gen/shortlink"

	"github.com/FinnTew/FinnEcommerce/src/internal/client/kitex_gen/shortlink/shortlinkservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() shortlinkservice.Client
	Service() string
	CreateShortLink(ctx context.Context, Req *shortlink.CreateShortLinkRequest, callOptions ...callopt.Option) (r *shortlink.CreateShortLinkResponse, err error)
	GetLongLink(ctx context.Context, Req *shortlink.GetLongLinkRequest, callOptions ...callopt.Option) (r *shortlink.GetLongLinkResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := shortlinkservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient shortlinkservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() shortlinkservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateShortLink(ctx context.Context, Req *shortlink.CreateShortLinkRequest, callOptions ...callopt.Option) (r *shortlink.CreateShortLinkResponse, err error) {
	return c.kitexClient.CreateShortLink(ctx, Req, callOptions...)
}

func (c *clientImpl) GetLongLink(ctx context.Context, Req *shortlink.GetLongLinkRequest, callOptions ...callopt.Option) (r *shortlink.GetLongLinkResponse, err error) {
	return c.kitexClient.GetLongLink(ctx, Req, callOptions...)
}
