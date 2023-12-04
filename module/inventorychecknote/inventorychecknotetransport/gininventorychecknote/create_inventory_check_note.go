package gininventorychecknote

import (
	"net/http"
	"qtk-store-api/common"
	"qtk-store-api/component/appctx"
	"qtk-store-api/component/generator"
	"qtk-store-api/middleware"
	"qtk-store-api/module/inventorychecknote/inventorychecknotebiz"
	"qtk-store-api/module/inventorychecknote/inventorychecknotemodel"
	"qtk-store-api/module/inventorychecknote/inventorychecknoterepo"
	"qtk-store-api/module/inventorychecknote/inventorychecknotestore"
	"qtk-store-api/module/inventorychecknotedetail/inventorychecknotedetailstore"
	"qtk-store-api/module/product/productstore"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Create inventory check note
// @Tags inventoryCheckNotes
// @Accept json
// @Produce json
// @Param inventoryCheckNote body inventorychecknotemodel.ReqCreateInventoryCheckNote true "inventory check note need to create"
// @Response 200 {object} inventorychecknotemodel.ResCreateInventoryCheckNote "inventory check note id"
// @Response 400 {object} common.AppError "error"
// @Router /inventoryCheckNotes [post]
func CreateInventoryCheckNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data inventorychecknotemodel.ReqCreateInventoryCheckNote

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreateBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		inventoryCheckNoteStore := inventorychecknotestore.NewSQLStore(db)
		inventoryCheckNoteDetailStore := inventorychecknotedetailstore.NewSQLStore(db)
		productStore := productstore.NewSQLStore(db)

		repo := inventorychecknoterepo.NewCreateInventoryCheckNoteRepo(
			inventoryCheckNoteStore,
			inventoryCheckNoteDetailStore,
			productStore,
		)

		gen := generator.NewShortIdGenerator()

		business := inventorychecknotebiz.NewCreateInventoryCheckNoteBiz(gen, repo, requester)

		if err := business.CreateInventoryCheckNote(c.Request.Context(), &data); err != nil {
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
