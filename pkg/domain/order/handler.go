package order

import "github.com/gin-gonic/gin"

// this handler will use GIN GONIC as http web framework
type OrderHandler interface {
	GetOrderByIDHdl(ctx *gin.Context)
	InsertOrderHdl(ctx *gin.Context)
}
