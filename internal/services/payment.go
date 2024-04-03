package services

import (
	"context"
	"google.golang.org/grpc"
	payments "parking-app-mobile-api-gateway/internal/payments/proto"
	"parking-app-mobile-api-gateway/internal/utils"
)

type PaymentService struct {
	URL string
}

func NewPaymentService() *PaymentService {
	url, _ := utils.GetEnvVar("PAYMENT_SERVICE_GRPC_URL")

	return &PaymentService{
		URL: url,
	}
}

func (p *PaymentService) GetPayments() *payments.GetAllPaymentsResponse {
	conn, err := grpc.Dial(p.URL, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := payments.NewPaymentsGrpcClient(conn)

	ctx := context.Background()
	r, err := c.GetAllPayments(ctx, &payments.GetAllPaymentsRequest{})
	if err != nil {
		panic(err)
	}

	return r
}

func (p *PaymentService) CreatePayment(createPaymentRequest *payments.CreatePaymentRequest) *payments.CreatePaymentResponse {
	conn, err := grpc.Dial(p.URL, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := payments.NewPaymentsGrpcClient(conn)

	ctx := context.Background()
	r, err := c.CreatePayment(ctx, createPaymentRequest)
	if err != nil {
		panic(err)
	}

	return r
}

func (p *PaymentService) GetPayment(id string) *payments.GetPaymentResponse {
	conn, err := grpc.Dial(p.URL, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := payments.NewPaymentsGrpcClient(conn)

	ctx := context.Background()
	r, err := c.GetPayment(ctx, &payments.GetPaymentRequest{Id: id})
	if err != nil {
		panic(err)
	}

	return r
}
