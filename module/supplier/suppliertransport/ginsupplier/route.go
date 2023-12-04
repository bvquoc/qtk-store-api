package ginsupplier

import (
	"qtk-store-api/component/appctx"
	"qtk-store-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	suppliers := router.Group("/suppliers", middleware.RequireAuth(appCtx))
	{
		suppliers.GET("", ListSupplier(appCtx))
		suppliers.POST("", CreateSupplier(appCtx))
		suppliers.GET("/:id/import_notes", SeeSupplierImportNote(appCtx))
		suppliers.GET("/:id/debts", SeeSupplierDebt(appCtx))
		suppliers.PATCH("/:id", UpdateInfoSupplier(appCtx))
		suppliers.POST("/:id/pay", PaySupplier(appCtx))
	}
}
