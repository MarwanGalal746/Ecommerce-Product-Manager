package productHandlers

import (
	"Ecommerce-Product-Manager/pkg/errs"
	"encoding/csv"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (productHandler ProductHandlers) Update(c *gin.Context) {
	reader := csv.NewReader(c.Request.Body)
	err := productHandler.Service.UpdateByCSV(reader)

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
}
