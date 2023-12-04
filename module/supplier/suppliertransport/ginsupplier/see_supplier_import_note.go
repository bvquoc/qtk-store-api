package ginsupplier

import (
	"net/http"
	"qtk-store-api/common"
	"qtk-store-api/component/appctx"
	"qtk-store-api/middleware"
	"qtk-store-api/module/importnote/importnotestore"
	"qtk-store-api/module/supplier/supplierbiz"
	"qtk-store-api/module/supplier/suppliermodel/filter"
	"qtk-store-api/module/supplier/supplierrepo"
	"qtk-store-api/module/supplier/supplierstore"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary See import notes of supplier
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path string true "supplier id"
// @Param page query common.Paging false "page"
// @Param filter query filter.SupplierDebtFilter false "filter"
// @Response 200 {object} suppliermodel.ResSeeDebtSupplier "supplier"
// @Response 400 {object} common.AppError "error"
// @Router /suppliers/{id}/import_notes [get]
func SeeSupplierImportNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var importSupplierFilter filter.SupplierImportFilter
		if err := c.ShouldBind(&importSupplierFilter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		importNoteStore := importnotestore.NewSQLStore(appCtx.GetMainDBConnection())
		supplierStore := supplierstore.NewSQLStore(appCtx.GetMainDBConnection())

		repo := supplierrepo.NewSeeSupplierImportNoteRepo(importNoteStore, supplierStore)
		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := supplierbiz.NewSeeSupplierImportNoteBiz(repo, requester)

		result, err := biz.SeeSupplierImportNote(
			c.Request.Context(), id, &importSupplierFilter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
