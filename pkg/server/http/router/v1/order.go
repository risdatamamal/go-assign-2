package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/risdatamamal/go-assign-2/config/gin"
	"github.com/risdatamamal/go-assign-2/pkg/server/http/router"
)

type OrderImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
}

func NewOrderRouter(ginEngine engine.HttpServer) router.Router {
	routerGroup := ginEngine.GetGin().Group("/v1/orders")
	return &OrderImpl{ginEngine: ginEngine, routerGroup: routerGroup}
}

func (o *OrderImpl) post() {
	// all path for post method are here
	o.routerGroup.POST("", func(ctx *gin.Context) {})
}

func (o *OrderImpl) Routers() {
	o.post()
}
