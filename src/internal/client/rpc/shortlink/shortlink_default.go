package shortlink

import (
	"context"
	shortlink "github.com/FinnTew/FinnEcommerce/src/internal/client/kitex_gen/shortlink"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func CreateShortLink(ctx context.Context, req *shortlink.CreateShortLinkRequest, callOptions ...callopt.Option) (resp *shortlink.CreateShortLinkResponse, err error) {
	resp, err = defaultClient.CreateShortLink(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateShortLink call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetLongLink(ctx context.Context, req *shortlink.GetLongLinkRequest, callOptions ...callopt.Option) (resp *shortlink.GetLongLinkResponse, err error) {
	resp, err = defaultClient.GetLongLink(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetLongLink call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
