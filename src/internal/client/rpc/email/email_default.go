package email

import (
	"context"
	email "github.com/FinnTew/FinnEcommerce/src/internal/client/kitex_gen/email"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SendEmail(ctx context.Context, req *email.SendEmailRequest, callOptions ...callopt.Option) (resp *email.SendEmailResponse, err error) {
	resp, err = defaultClient.SendEmail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SendEmail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
