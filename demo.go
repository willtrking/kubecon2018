package main

import (
	"context"
	"fmt"
	"github.com/willtrking/kubecon2018/proto"
	"math/rand"
	"time"
)

var productIds = []string{
	"A5FF322A-91B6-44A6-923D-23496EAE6F25",
	"32F3ACD3-18AA-447E-8ADE-7412F2FE584F",
	"C9684D5A-9800-4FED-8FE5-45B64F402A29",
}

func demoRun(client proto.CartClient, numberOk, numberErr int) {

	rand.Seed(time.Now().UnixNano())

	var toAdd []*proto.AddRequest

	if numberOk > 0 {
		for {
			for _, id := range productIds {
				toAdd = append(toAdd, &proto.AddRequest{
					ProductId:id,
					AccountId: "real",
				})
				if len(toAdd) == numberOk {
					break
				}
			}

			if len(toAdd) >= numberOk {
				break
			}

		}
	}

	for i := 0; i < numberErr; i++ {

		toAdd = append(toAdd, &proto.AddRequest{
			ProductId:"A5FF322A-91B6-44A6-923D-23496EAE6F25",
			AccountId: "fake",
		})

	}


	rand.Shuffle(len(toAdd), func(i, j int) {
		toAdd[i], toAdd[j] = toAdd[j], toAdd[i]
	})

	for _, ta := range toAdd {
		fmt.Println(client.AddToCart(context.Background(), ta))
	}

}
