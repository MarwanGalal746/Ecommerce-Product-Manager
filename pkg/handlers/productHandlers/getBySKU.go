package productHandlers

import (
	"Ecommerce-Product-Manager/pkg/errs"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (productHandler ProductHandlers) Get(c *gin.Context) {
	product, err := productHandler.Service.Get(c.Param("sku"))
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	} else if err != nil {
		log.Println(err.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(product)
}
