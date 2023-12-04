package authortransport

import (
	"qtk-store-api/component/appctx"
	"qtk-store-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	authors := router.Group("/authors", middleware.RequireAuth(appCtx))
	{
		authors.GET("", ListAuthor(appCtx))
		authors.POST("", CreateAuthor(appCtx))
	}
}
