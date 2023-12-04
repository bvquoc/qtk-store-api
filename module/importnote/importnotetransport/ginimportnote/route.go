package ginimportnote

import (
	"qtk-store-api/component/appctx"
	"qtk-store-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appCtx appctx.AppContext) {
	importNotes := router.Group("/importNotes", middleware.RequireAuth(appCtx))
	{
		importNotes.GET("", ListImportNote(appCtx))
		importNotes.GET("/:id", SeeDetailImportNote(appCtx))
		importNotes.POST("", CreateImportNote(appCtx))
		importNotes.PATCH("/:id", ChangeStatusImportNote(appCtx))
	}
}
