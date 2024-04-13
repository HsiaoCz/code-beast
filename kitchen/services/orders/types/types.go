package types

import (
	"context"

	"github.com/HsiaoCz/code-beast/kitchen/services/common/genproto/orders"
)

type OrderServicer interface {
	CreateOrder(context.Context, *orders.Order) error
}
