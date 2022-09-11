package productHandlers

import (
	"Ecommerce-Product-Manager/pkg/config"
	"Ecommerce-Product-Manager/pkg/errs"
	"encoding/csv"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (productHandler ProductHandlers) Update(c *gin.Context) {
	reader := csv.NewReader(c.Request.Body)
	err := productHandler.Service.UpdateByCSV(reader)
	config.Logger.Info("extracting CSV file from request")
	if err != nil && err.Error() == errs.ErrDb.Error() {
		config.Logger.Error(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		config.Logger.Debug("API response status code is " + http.StatusText(http.StatusInternalServerError))
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return

	} else if err != nil {
		config.Logger.Error(err.Error())
		config.Logger.Debug("API response status code is " + http.StatusText(http.StatusInternalServerError))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	config.Logger.Debug("API response status code is " + http.StatusText(http.StatusOK))
	json.NewEncoder(c.Writer).Encode(errs.NewResponse("products stocks has been updated successfully",
		http.StatusOK))
}
