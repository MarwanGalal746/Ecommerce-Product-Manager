package errs

import "errors"

var (
	ErrDb                                        = errors.New("unexpected database error")
	ErrStockFromProductNotAvailableInThisCountry = errors.New("stock from this product is not available in this country")
	ErrProductDoesNotExist                       = errors.New("product does not exist")
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewResponse(message string, status int) *Response {
	return &Response{Message: message, Status: status}
}
