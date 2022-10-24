package order

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/risdatamamal/go-assign-2/pkg/domain/message"
	"github.com/risdatamamal/go-assign-2/pkg/domain/order"
)

type OrderHdlImpl struct {
	orderUsecase order.OrderUsecase
}

func NewOrderHandler(orderUsecase order.OrderUsecase) order.OrderHandler {
	return &OrderHdlImpl{orderUsecase: orderUsecase}
}

func (u *OrderHdlImpl) GetOrderByIDHdl(ctx *gin.Context) {

}

// Insert New Order
// @Summary this api will insert order
// @Schemes
// @Description insert new order
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} Order
// @Router /v1/orders [post]
func (o *OrderHdlImpl) InsertOrderHdl(ctx *gin.Context) {
	// dengan menggunakan context gin,
	// kita bisa langsung mendapatkan value dan set value dari function didalam context tsb

	// set value KEY2 VALUE2
	ctx.Set("KEY2", "VALUE2")

	// get value from context
	key2 := ctx.Value("KEY2")
	fmt.Println(key2)

	log.Printf("%T - InsertOrderHdl is invoked]\n", o)
	defer log.Printf("%T - InsertOrderHdl executed\n", o)

	// binding / mendapatkan body payload dari request
	log.Println("binding body payload from request")
	var order order.Order
	if err := ctx.ShouldBind(&order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "failed to bind payload",
		})
		return
	}
	// check apakah id kosong atau tidak: kalau kosong lempar BAD_REQUEST
	log.Println("check id from request")
	if strconv.FormatUint(order.OrderID, int(order.OrderID)) == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "id should not be empty",
		})
		return
	}
	// call service/usecase untuk menginsert data
	log.Println("calling insert service usecase")

	result, err := o.orderUsecase.InsertOrderSvc(ctx, order)
	if err != nil {
		switch err.Error() {
		case "BAD_REQUEST":
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Code:  80,
				Error: "invalid processing payload",
			})
			return
		case "INTERNAL_SERVER_ERROR":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
				Code:  99,
				Error: "something went wrong",
			})
			return
		}
	}
	// response result for the order if success
	ctx.JSONP(http.StatusOK, message.Response{
		Code:    0,
		Message: "success insert order",
		Data:    result,
	})
}
