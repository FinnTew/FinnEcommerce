package service

import (
	"context"
	payment "github.com/FinnTew/FinnEcommerce/src/internal/service/payment/kitex_gen/payment"
	"testing"
)

func TestCancelPayment_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCancelPaymentService(ctx)
	// init req and assert value

	req := &payment.CancelPaymentReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
