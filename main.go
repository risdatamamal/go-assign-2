package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	engine "github.com/risdatamamal/go-assign-2/config/gin"
	"github.com/risdatamamal/go-assign-2/config/postgres"
	docs "github.com/risdatamamal/go-assign-2/docs"
	"github.com/risdatamamal/go-assign-2/pkg/domain/message"
	orderrepo "github.com/risdatamamal/go-assign-2/pkg/repository/order"
	orderhandler "github.com/risdatamamal/go-assign-2/pkg/server/http/handler/order"
	v1 "github.com/risdatamamal/go-assign-2/pkg/server/http/router/v1"
	orderusecase "github.com/risdatamamal/go-assign-2/pkg/usecase/order"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

// comment dalam go
// untuk beberapa CODE GENERATOR -> tools yang digunakan untuk
// membuat code template di dalam project GO
// ex: swaggo, mockgen, dll
// untuk beberapa tools generator, tools akan membaca comment
// yang memiliki annotation

// @title UserOrder API
// @version 1.0
// @description This is api for create, read, update, delete order and item order
// @termOfService https://swagger.io/terms
// @contact.name Tamam API Support
// @host localhost:8080
// @BasePath /
func main() {
	// generate postgres config and connect to postgres
	// this postgres client, will be used in repository layer
	postgresCln := postgres.NewPostgresConnection(postgres.Config{
		Host:         "localhost",
		Port:         "5432",
		User:         "postgres",
		Password:     "postgresAdmin",
		DatabaseName: "postgres",
	})

	// gin engine
	ginEngine := engine.NewGinHttp(engine.Config{
		Port: ":8080",
	})

	// setiap request yang datang ke API ini,
	// dia akan melalui gin.Recovery dan gin.Logger
	// .USE disini, adalah cara untuk memasukkan middleware juga
	ginEngine.GetGin().Use(
		gin.Recovery(),
		gin.Logger())

	startTime := time.Now()
	ginEngine.GetGin().GET("/", func(ctx *gin.Context) {
		// secara default map jika di return dalam
		// response API, dia akan menjadi JSON
		respMap := map[string]any{
			"code":       0,
			"message":    "server up and running",
			"start_time": startTime,
		}

		// golang memiliki json package
		// json package bisa mentranslasikan
		// map menjadi suatu struct
		// nb: struct harus memiliki tag/annotation JSON
		var respStruct message.Response

		// marshal -> mengubah json/struct/map menjadi
		// array of byte atau bisa kita translatekan menjadi string
		// dengan format JSON
		resByte, err := json.Marshal(respMap)
		if err != nil {
			panic(err)
		}
		// unmarshal -> translasikan string/[]byte dengan format JSON
		// menjadi map/struct dengan tag/annotation json
		err = json.Unmarshal(resByte, &respStruct)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, respStruct)
	})

	docs.SwaggerInfo.BasePath = "/v1"
	ginEngine.GetGin().GET("/swagger/*any", ginswagger.
		WrapHandler(swaggerfiles.Handler))

	// generate order repository
	orderRepo := orderrepo.NewOrderRepo(postgresCln)
	// initiate order use case
	orderUsecase := orderusecase.NewOrderUsecase(orderRepo)
	// initiate handler
	useHandler := orderhandler.NewOrderHandler(orderUsecase)
	// initiate router
	v1.NewOrderRouter(ginEngine, useHandler).Routers()
	v1.NewOrderRouter(ginEngine).Routers()

	// running the service
	ginEngine.Serve()
}
