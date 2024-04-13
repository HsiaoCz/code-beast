package handlers

import (
	"context"

	"github.com/HsiaoCz/code-beast/kitchen/services/common/genproto/orders"
	"github.com/HsiaoCz/code-beast/kitchen/services/orders/types"
)

type OrderGrpcHandler struct {
	orderService types.OrderServicer
	orders.UnimplementedOrderServiceServer
}

func NewOrderGrpcHandler(orderService types.OrderServicer) *OrderGrpcHandler {
	return &OrderGrpcHandler{
		orderService: orderService,
	}
}

func (og *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	return &orders.CreateOrderResponse{}, nil
}

func (og *OrderGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrdersResponse, error) {
	return &orders.GetOrdersResponse{}, nil
}
