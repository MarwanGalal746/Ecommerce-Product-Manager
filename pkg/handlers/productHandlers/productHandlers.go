package productRepository

import "github.com/gin-gonic/gin"

type ProductHandlers interface {
	Get(c *gin.Context)
	ConsumeStock(c *gin.Context)
	Update(c *gin.Context)
}
