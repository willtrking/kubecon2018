package account

import (
	"time"
	"github.com/willtrking/kubecon2018/proto"
	"context"
	"math/rand"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"
)

type AccountGRPC struct {
}

func randomLatency() {
	// Introduce latency, random between 25 and 100 ms
	time.Sleep(time.Duration(rand.Intn(75)+25) * time.Millisecond)
}

func (c *AccountGRPC) GetAccount(ctx context.Context, req *proto.AccountReq) (*proto.Account, error) {

	randomLatency()

	if req.AccountId == "real" {
		return &proto.Account{
			AccountId:"real",
		}, nil
	}

	return nil, status.Error(codes.InvalidArgument, "Invalid account ID")

}
