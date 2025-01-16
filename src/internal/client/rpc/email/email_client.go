package email

import (
	"context"
	email "github.com/FinnTew/FinnEcommerce/src/internal/client/kitex_gen/email"

	"github.com/FinnTew/FinnEcommerce/src/internal/client/kitex_gen/email/emailservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() emailservice.Client
	Service() string
	SendEmail(ctx context.Context, Req *email.SendEmailRequest, callOptions ...callopt.Option) (r *email.SendEmailResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := emailservice.NewClient(dstService, opts...)
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
	kitexClient emailservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() emailservice.Client {
	return c.kitexClient
}

func (c *clientImpl) SendEmail(ctx context.Context, Req *email.SendEmailRequest, callOptions ...callopt.Option) (r *email.SendEmailResponse, err error) {
	return c.kitexClient.SendEmail(ctx, Req, callOptions...)
}
