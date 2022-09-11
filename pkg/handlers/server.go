package handlers

import (
	"Ecommerce-Product-Manager/pkg/domain/repositories/productRepository/defaultProductRepository"
	"Ecommerce-Product-Manager/pkg/handlers/productHandlers"
	"Ecommerce-Product-Manager/pkg/service/productService/defaultProductService"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	productHandler := productHandlers.ProductHandlers{
		Service: defaultProductService.NewDefaultProductService(defaultProductRepository.DefaultProductRepositoryDb{})}

	router.GET("/get/:sku", productHandler.Get)
	router.POST("/consume/:sku/:country/:stock", productHandler.ConsumeStock)
	go router.POST("/update", productHandler.Update)

	router.Run(":8888")

}
