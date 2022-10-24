package order

import "context"

type OrderRepo interface {
	GetOrderByID(ctx context.Context, id uint64) (order Order, err error)
	InsertOrder(ctx context.Context, order *Order) (err error)
}
