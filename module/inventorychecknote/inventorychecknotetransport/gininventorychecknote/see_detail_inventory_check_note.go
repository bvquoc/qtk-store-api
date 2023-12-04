package gininventorychecknote

import (
	"net/http"
	"qtk-store-api/common"
	"qtk-store-api/component/appctx"
	"qtk-store-api/middleware"
	"qtk-store-api/module/inventorychecknote/inventorychecknotebiz"
	"qtk-store-api/module/inventorychecknote/inventorychecknoterepo"
	"qtk-store-api/module/inventorychecknote/inventorychecknotestore"
	"qtk-store-api/module/inventorychecknotedetail/inventorychecknotedetailstore"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See detail inventory check note
// @Tags inventoryCheckNotes
// @Accept json
// @Produce json
// @Param id path string true "inventory check note id"
// @Param page query common.Paging false "page"
// @Response 200 {object} inventorychecknotemodel.ResSeeDetailInventoryCheckNote "inventory check note"
// @Response 400 {object} common.AppError "error"
// @Router /inventoryCheckNotes/{id} [get]
func SeeDetailInventoryCheckNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		inventoryCheckNoteId := c.Param("id")

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		inventoryCheckNoteStore :=
			inventorychecknotestore.NewSQLStore(appCtx.GetMainDBConnection())
		inventoryCheckNoteDetailStore :=
			inventorychecknotedetailstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := inventorychecknoterepo.NewSeeDetailInventoryCheckNoteRepo(
			inventoryCheckNoteStore, inventoryCheckNoteDetailStore)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := inventorychecknotebiz.NewSeeDetailImportNoteBiz(repo, requester)

		result, err := biz.SeeDetailInventoryCheckNote(
			c.Request.Context(), inventoryCheckNoteId, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, &paging, nil))
	}
}
