package ginimportnote

import (
	"net/http"
	"qtk-store-api/common"
	"qtk-store-api/component/appctx"
	"qtk-store-api/component/generator"
	"qtk-store-api/middleware"
	"qtk-store-api/module/importnote/importnotebiz"
	"qtk-store-api/module/importnote/importnotemodel"
	"qtk-store-api/module/importnote/importnoterepo"
	"qtk-store-api/module/importnote/importnotestore"
	"qtk-store-api/module/importnotedetail/importnotedetailstore"
	"qtk-store-api/module/product/productstore"
	"qtk-store-api/module/supplier/supplierstore"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Create import note
// @Tags importNotes
// @Accept json
// @Produce json
// @Param importNote body importnotemodel.ReqCreateImportNote true "import note need to create"
// @Response 200 {object} importnotemodel.ResCreateImportNote "import note id"
// @Response 400 {object} common.AppError "error"
// @Router /importNotes [post]
func CreateImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data importnotemodel.ReqCreateImportNote

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreateBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		importNoteStore := importnotestore.NewSQLStore(db)
		importNoteDetailStore := importnotedetailstore.NewSQLStore(db)
		productStore := productstore.NewSQLStore(db)
		supplierStore := supplierstore.NewSQLStore(db)

		repo := importnoterepo.NewCreateImportNoteRepo(
			importNoteStore,
			importNoteDetailStore,
			productStore,
			supplierStore,
		)

		gen := generator.NewShortIdGenerator()

		business := importnotebiz.NewCreateImportNoteBiz(gen, repo, requester)

		if err := business.CreateImportNote(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
