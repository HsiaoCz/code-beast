package handlers

import (
	"context"

	"github.com/HsiaoCz/code-beast/kitchen/services/common/genproto/orders"
	"github.com/HsiaoCz/code-beast/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrderGrpcHandler struct {
	orderService types.OrderServicer
	orders.UnimplementedOrderServiceServer
}

func NewOrderGrpcHandler(grpc *grpc.Server, orderService types.OrderServicer) {
	grpcHandler := &OrderGrpcHandler{
		orderService: orderService,
	}
	// register the OrderServiceServer
	orders.RegisterOrderServiceServer(grpc, grpcHandler)
}

func (og *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
	}
	if err := og.orderService.CreateOrder(ctx, order); err != nil {
		return nil, err
	}
	return &orders.CreateOrderResponse{
		Status: "success",
	}, nil
}

func (og *OrderGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrdersResponse, error) {
	return &orders.GetOrdersResponse{}, nil
}
