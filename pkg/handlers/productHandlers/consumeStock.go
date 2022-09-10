package productHandlers

import (
	"Ecommerce-Product-Manager/pkg/errs"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (productHandler ProductHandlers) ConsumeStock(c *gin.Context) {
	err := productHandler.Service.ConsumeStock(c.Param("sku"),
		c.Param("country"), c.Param("stock"))

	if err != nil && err.Error() == errs.ErrDb.Error() {

		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return

	} else if err != nil && err.Error() == errs.ErrProductNotAvailableInThisCountry.Error() {

		log.Println(errs.ErrProductNotAvailableInThisCountry.Error())
		c.Writer.WriteHeader(http.StatusOK)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return

	} else if err != nil {

		log.Println(err.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return

	}

	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(errs.NewResponse("product stock has been updated successfully",
		http.StatusOK))
}
