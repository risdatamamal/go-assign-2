package order

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/risdatamamal/go-assign-2/config/postgres"
	"github.com/risdatamamal/go-assign-2/pkg/domain/order"
)

type OrderRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewOrderRepo(pgCln postgres.PostgresClient) order.OrderRepo {
	return &OrderRepoImpl{pgCln: pgCln}
}

func (o *OrderRepoImpl) GetOrderByID(ctx context.Context, id uint64) (result order.Order, err error) {
	log.Printf("%T - GetOrderByID is invoked]\n", o)
	defer log.Printf("%T - GetOrderByID executed\n", o)
	// get gorm client first
	db := o.pgCln.GetClient().WithContext(ctx)
	// insert new order
	db.Model(&order.Order{}).
		Where("id = ?", id).
		Find(&result)
	//check error
	if err = db.Error; err != nil {
		log.Printf("error when getting order with id %v\n",
			id)
	}
	return result, err
}

func (o *OrderRepoImpl) InsertOrder(ctx context.Context, insertedOrder *order.Order) (err error) {
	log.Printf("%T - InsertOrder is invoked]\n", o)
	defer log.Printf("%T - InsertOrder executed\n", o)

	dl, _ := ctx.Deadline()
	if time.Now().After(dl) {
		// context reach deadline
		return errors.New("context canceled by deadline")
	}

	// only valid if context include KEY1
	key1 := ctx.Value("KEY1")
	if key1 == nil {
		// context doesn't contain KEY1
		return errors.New("context invalid")
	}

	key2 := ctx.Value("KEY2")
	if key2 == nil {
		// context doesn't contain KEY1
		return errors.New("context invalid")
	}

	// get gorm client first
	db := o.pgCln.GetClient().WithContext(ctx)
	// insert new order
	db.Model(&order.Order{}).
		Create(&insertedOrder)
	//check error
	if err = db.Error; err != nil {
		log.Printf("error when inserting order with id %v\n",
			insertedOrder.OrderID)
	}
	return err
}
