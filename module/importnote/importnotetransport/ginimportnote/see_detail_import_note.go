package ginimportnote

import (
	"net/http"
	"qtk-store-api/common"
	"qtk-store-api/component/appctx"
	"qtk-store-api/middleware"
	"qtk-store-api/module/importnote/importnotebiz"
	"qtk-store-api/module/importnote/importnoterepo"
	"qtk-store-api/module/importnote/importnotestore"
	"qtk-store-api/module/importnotedetail/importnotedetailstore"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See detail import note
// @Tags importNotes
// @Accept json
// @Produce json
// @Param id path string true "import note id"
// @Param page query common.Paging false "page"
// @Response 200 {object} importnotemodel.ResSeeDetailImportNote "import note"
// @Response 400 {object} common.AppError "error"
// @Router /importNotes/{id} [get]
func SeeDetailImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		importNoteId := c.Param("id")

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		importNoteDetailStore := importnotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		importNoteStore := importnotestore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := importnoterepo.NewSeeDetailImportNoteRepo(
			importNoteStore, importNoteDetailStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := importnotebiz.NewSeeDetailImportNoteBiz(
			repo, requester)

		result, err := biz.SeeDetailImportNote(c.Request.Context(), importNoteId, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, &paging, nil))
	}
}
