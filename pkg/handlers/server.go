package handlers

import (
	"Ecommerce-Product-Manager/pkg/config/configSQL"
	"Ecommerce-Product-Manager/pkg/domain/repositories/productRepository/defaultProductRepository"
	"Ecommerce-Product-Manager/pkg/handlers/productHandlers/defaultProductHandlers"
	"Ecommerce-Product-Manager/pkg/service/productService/defaultProductService"
	"github.com/gin-gonic/gin"
)

type Server struct {
	configSQLclient configSQL.ConfigPgSQL
	productHandler  defaultProductHandlers.DefaultProductHandlers
}

func (server *Server) Start() {
	server.configSQLclient = configSQL.ConfigPgSQL{}
	SQLdb, err := server.configSQLclient.Config()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	server.productHandler = defaultProductHandlers.NewDefaultProductHandlers(
		defaultProductService.NewDefaultProductService(
			defaultProductRepository.NewDefaultProductRepositoryDb(SQLdb)))

	router.GET("/get/:sku", server.productHandler.Get)
	router.POST("/consume/:sku/:country/:stock", server.productHandler.ConsumeStock)
	go router.POST("/update", server.productHandler.Update)

	router.Run(":8888")

}
