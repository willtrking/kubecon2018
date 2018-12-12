package merchandise

import (
	"context"
	"github.com/willtrking/kubecon2018/proto"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"
	"math/rand"
	"time"
)

type MerchandiseGRPC struct {
}

func randomLatency() {
	// Introduce latency, random between 25 and 100 ms
	time.Sleep(time.Duration(rand.Intn(75)+25) * time.Millisecond)
}

func (m *MerchandiseGRPC) GetProductByID(ctx context.Context, req *proto.ProductRequest) (*proto.Product, error) {

	randomLatency()

	if productName, hasProduct := products[req.Id]; hasProduct {

		return &proto.Product{
			Id:   req.Id,
			Name: productName,
		}, nil
	}

	return nil, status.Error(codes.InvalidArgument, "Invalid product ID")

}
