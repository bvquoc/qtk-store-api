package categorytransport

import (
	"qtk-store-api/component/appctx"
	"qtk-store-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	categories := router.Group("/categories", middleware.RequireAuth(appCtx))
	{
		categories.GET("", ListCategory(appCtx))
		categories.POST("", CreateCategory(appCtx))
	}
}
