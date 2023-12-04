package ginimportnote

import (
	"net/http"
	"qtk-store-api/common"
	"qtk-store-api/component/appctx"
	"qtk-store-api/middleware"
	"qtk-store-api/module/importnote/importnotebiz"
	"qtk-store-api/module/importnote/importnotemodel"
	"qtk-store-api/module/importnote/importnoterepo"
	"qtk-store-api/module/importnote/importnotestore"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary List import note
// @Tags importNotes
// @Accept json
// @Produce json
// @Param page query common.Paging false "page"
// @Param filter query importnotemodel.Filter false "filter"
// @Response 200 {object} importnotemodel.ResListImportNote "list import note"
// @Response 400 {object} common.AppError "error"
// @Router /importNotes [get]
func ListImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter importnotemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := importnotestore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := importnoterepo.NewListImportNoteRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := importnotebiz.NewListImportNoteBiz(repo, requester)

		result, err := biz.ListImportNote(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
