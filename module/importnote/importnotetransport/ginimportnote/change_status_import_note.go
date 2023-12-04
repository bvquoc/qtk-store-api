package ginimportnote

import (
	"errors"
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
	"qtk-store-api/module/supplierdebt/supplierdebtstore"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Change status import note
// @Tags importNotes
// @Accept json
// @Produce json
// @Param id path string true "import note id"
// @Param importNote body importnotemodel.ReqUpdateImportNote true "status need to update of import note"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /importNotes/{id} [patch]
func ChangeStatusImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idImportNote := c.Param("id")
		if idImportNote == "" {
			panic(common.ErrInvalidRequest(errors.New("param id not exist")))
		}

		var data importnotemodel.ReqUpdateImportNote

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CloseBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		importNoteStore := importnotestore.NewSQLStore(db)
		importNoteDetailStore := importnotedetailstore.NewSQLStore(db)
		productStore := productstore.NewSQLStore(db)
		supplierStore := supplierstore.NewSQLStore(db)
		supplierDebtStore := supplierdebtstore.NewSQLStore(db)

		repo := importnoterepo.NewChangeStatusImportNoteRepo(
			importNoteStore,
			importNoteDetailStore,
			productStore,
			supplierStore,
			supplierDebtStore,
		)

		gen := generator.NewShortIdGenerator()

		business := importnotebiz.NewChangeStatusImportNoteBiz(gen, repo, requester)

		if err := business.ChangeStatusImportNote(
			c.Request.Context(),
			idImportNote,
			&data,
		); err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
