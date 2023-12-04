package producttransport

import (
	"qtk-store-api/component/appctx"
	"qtk-store-api/middleware"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Get all products
// @Tags products
// @Accept json
// @Produce json
// @Response 200 {object} productmodel.Product
// @Router /products [get]
func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	products := router.Group("/products", middleware.RequireAuth(appCtx))
	{
		products.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "get all products",
			})
		})
		//products.POST("", CreateProduct(appCtx))
	}
}
