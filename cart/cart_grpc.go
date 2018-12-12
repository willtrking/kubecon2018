package cart

import (
	"context"
	"github.com/willtrking/kubecon2018/proto"
	"math/rand"
	"time"
)

type CartGRPC struct {
	MerchandiseClient proto.MerchandiseClient
	AccountsClient proto.AccountsClient
}

func randomLatency() {
	// Introduce latency, random between 25 and 100 ms
	time.Sleep(time.Duration(rand.Intn(75)+25) * time.Millisecond)
}

func (c *CartGRPC) ViewCart(ctx context.Context, req *proto.CartRequest) (*proto.UserCart, error) {

	randomLatency()

	_, err := c.AccountsClient.GetAccount(ctx,&proto.AccountReq{
		AccountId:req.AccountId,
	})

	if err != nil {
		return nil, err
	}

	return &proto.UserCart{
		Products: getCart(req.AccountId),
	}, nil
}

func (c *CartGRPC) AddToCart(ctx context.Context, req *proto.AddRequest) (*proto.Product, error) {

	randomLatency()

	_, err := c.AccountsClient.GetAccount(ctx,&proto.AccountReq{
		AccountId:req.AccountId,
	})

	if err != nil {
		return nil, err
	}

	product, err := c.MerchandiseClient.GetProductByID(ctx, &proto.ProductRequest{
		Id: req.ProductId,
	})

	if err != nil {
		return nil, err
	}

	addToCart(req.AccountId, product)

	return product, nil
}
