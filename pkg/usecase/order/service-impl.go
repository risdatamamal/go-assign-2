package order

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/risdatamamal/go-assign-2/pkg/domain/order"
)

type OrderUsecaseImpl struct {
	orderRepo order.OrderRepo
}

func NewOrderUsecase(orderRepo order.OrderRepo) order.OrderUsecase {
	return &OrderUsecaseImpl{orderRepo: orderRepo}
}

// GetOrderByIDSvc implements order.OrderUsecase
func (o *OrderUsecaseImpl) GetOrderByIDSvc(ctx context.Context, id uint64) (result order.Order, err error) {
	log.Printf("%T - GetOrderByID is invoked]\n", o)
	defer log.Printf("%T - GetOrderByID executed\n", o)
	// get order from repository (database)
	log.Println("getting order from order repository")
	result, err = o.orderRepo.GetOrderByID(ctx, id)
	if err != nil {
		// ini berarti ada yang salah dengan connection di database
		log.Println("error when fetching data from database: " + err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")
		return result, err
	}
	// check user id > 0 ?
	log.Println("checking order id")
	if result.OrderID <= 0 {
		// kalau tidak berarti order not found
		log.Println("order is not found: " + strconv.FormatUint(id, int(id)))
		err = errors.New("NOT_FOUND")
		return result, err
	}
	return result, err
}

// InsertOrderSvc implements order.OrderUsecase
func (o *OrderUsecaseImpl) InsertOrderSvc(ctx context.Context, input order.Order) (result order.Order, err error) {
	log.Printf("%T - InsertOrderSvc is invoked]\n", o)
	defer log.Printf("%T - InsertOrderSvc executed\n", o)

	// function ini tidak dijalankan lebih dari 2 detik
	ctx, cancel := context.WithTimeout(ctx, 1000*time.Second)
	defer cancel()

	// set value in context
	ctx = context.WithValue(ctx, "KEY1", "VALUE1")

	// get order for input id first
	ordCheck, err := o.GetOrderByIDSvc(ctx, input.OrderID)

	// check order is exist or not
	if err == nil {
		// order found
		log.Printf("order has been registered with id: %v\n", ordCheck.OrderID)
		err = errors.New("BAD_REQUEST")
		return result, err
	}
	// internal server error condition
	if err.Error() != "NOT_FOUND" {
		// internal server error
		log.Println("got error when checking user from database")
		return result, err
	}
	// valid condition: NOT_FOUND
	log.Println("insert order to database process")
	if err = o.orderRepo.InsertOrder(ctx, &input); err != nil {
		log.Printf("error when inserting order:%v\n", err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")
	}
	return input, err
}
