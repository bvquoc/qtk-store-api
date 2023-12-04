package ginuser

import (
	"qtk-store-api/component/appctx"
	"qtk-store-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	router.POST("/login", Login(appCtx))
	users := router.Group("/users", middleware.RequireAuth(appCtx))
	{
		users.POST("", CreateUser(appCtx))
	}
}
