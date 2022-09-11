package productHandlers

import (
	"Ecommerce-Product-Manager/pkg/config"
	"Ecommerce-Product-Manager/pkg/errs"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (productHandler ProductHandlers) ConsumeStock(c *gin.Context) {
	err := productHandler.Service.ConsumeStock(c.Param("sku"),
		c.Param("country"), c.Param("stock"))
	config.Logger.Info("extracting product SKU, country which product should be " +
		"consumed and and the amount should be consumed from stock from request")

	if err != nil && err.Error() == errs.ErrDb.Error() {

		config.Logger.Error(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		config.Logger.Debug("API response status code is " + http.StatusText(http.StatusInternalServerError))
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return

	} else if err != nil && err.Error() == errs.ErrStockFromProductNotAvailableInThisCountry.Error() {
		config.Logger.Error(errs.ErrStockFromProductNotAvailableInThisCountry.Error())
		config.Logger.Debug("API response status code is " + http.StatusText(http.StatusOK))
		c.Writer.WriteHeader(http.StatusOK)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrStockFromProductNotAvailableInThisCountry.Error(), http.StatusInternalServerError))
		return

	} else if err != nil {
		config.Logger.Error(err.Error())
		config.Logger.Debug("API response status code is " + http.StatusText(http.StatusInternalServerError))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(err.Error(), http.StatusInternalServerError))
		return

	}
	config.Logger.Debug("API response status code is " + http.StatusText(http.StatusOK))
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(errs.NewResponse("product stock has been consumed successfully",
		http.StatusOK))
}
