package defaultProductHandlers

import (
	"Ecommerce-Product-Manager/pkg/service/productService"
)

type DefaultProductHandlers struct {
	Service productService.ProductService
}

func NewDefaultProductHandlers(Service productService.ProductService) DefaultProductHandlers {
	return DefaultProductHandlers{Service: Service}
}
