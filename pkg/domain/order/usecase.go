package order

import "context"

type OrderUsecase interface {
	GetOrderByIDSvc(ctx context.Context, id uint64) (result Order, err error)
	InsertOrderSvc(ctx context.Context, input Order) (result Order, err error)
}
