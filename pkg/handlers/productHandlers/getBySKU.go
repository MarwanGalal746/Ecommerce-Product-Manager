package productHandlers

import (
	"Ecommerce-Product-Manager/pkg/config"
	"Ecommerce-Product-Manager/pkg/errs"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (productHandler ProductHandlers) Get(c *gin.Context) {
	product, err := productHandler.Service.Get(c.Param("sku"))
	config.Logger.Info("extracting product sku from request")
	if err != nil && err.Error() == errs.ErrDb.Error() {
		config.Logger.Error(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		config.Logger.Debug("API response status code is " + http.StatusText(http.StatusInternalServerError))
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	} else if err.Error() == errs.ErrProductDoesNotExist.Error() {
		config.Logger.Error(errs.ErrProductDoesNotExist.Error())
		c.Writer.WriteHeader(http.StatusOK)
		config.Logger.Debug("API response status code is " + http.StatusText(http.StatusOK))
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrProductDoesNotExist.Error(), http.StatusOK))
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
	json.NewEncoder(c.Writer).Encode(product)
}
