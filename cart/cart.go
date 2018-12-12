package cart

import (
	"github.com/willtrking/kubecon2018/proto"
	"sync"
)

var lock sync.RWMutex
var cartStorage = map[string][]*proto.Product{}

func addToCart(accountId string, product *proto.Product) {
	lock.Lock()
	defer lock.Unlock()
	cartStorage[accountId] = append(cartStorage[accountId], product)

}

func getCart(accountId string) []*proto.Product {
	lock.RLock()
	defer lock.RUnlock()

	if products, hasCart := cartStorage[accountId]; hasCart {
		return products
	}

	return nil
}
