package productHandlers

import (
	"Ecommerce-Product-Manager/pkg/domain/models"
	"Ecommerce-Product-Manager/pkg/errs"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (productHandler ProductHandlers) Get(c *gin.Context) {
	var product models.Product
	sku := c.Param("sku")
	product.SKU = sku
	product, err := productHandler.Service.Get(product)
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(product)
}
