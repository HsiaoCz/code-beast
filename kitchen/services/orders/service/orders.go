package service

import (
	"context"
	"fmt"

	"github.com/HsiaoCz/code-beast/kitchen/services/common/genproto/orders"
)

var ordersMap = make([]*orders.Order, 0)

type OrderService struct {
	// store
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersMap = append(ordersMap, order)
	fmt.Println(ordersMap)
	return nil
}
