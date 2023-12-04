package gininventorychecknote

import (
	"qtk-store-api/component/appctx"
	"qtk-store-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	inventoryCheckNotes := router.Group("/inventoryCheckNotes", middleware.RequireAuth(appCtx))
	{
		inventoryCheckNotes.GET("", ListInventoryCheckNote(appCtx))
		inventoryCheckNotes.GET("/:id", SeeDetailInventoryCheckNote(appCtx))
		inventoryCheckNotes.POST("", CreateInventoryCheckNote(appCtx))
	}
}
