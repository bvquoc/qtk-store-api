package ginsupplier

import (
	"net/http"
	"qtk-store-api/common"
	"qtk-store-api/component/appctx"
	"qtk-store-api/middleware"
	"qtk-store-api/module/supplier/supplierbiz"
	"qtk-store-api/module/supplier/suppliermodel/filter"
	"qtk-store-api/module/supplier/supplierrepo"
	"qtk-store-api/module/supplier/supplierstore"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary List supplier
// @Tags suppliers
// @Accept json
// @Produce json
// @Param page query common.Paging false "page"
// @Param filter query filter.Filter false "filter"
// @Response 200 {object} suppliermodel.ResListSupplier "list supplier"
// @Response 400 {object} common.AppError "error"
// @Router /suppliers [get]
func ListSupplier(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var supplierFilter filter.Filter
		if err := c.ShouldBind(&supplierFilter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := supplierstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := supplierrepo.NewListSupplierRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := supplierbiz.NewListSupplierRepo(repo, requester)

		result, err := biz.ListSupplier(c.Request.Context(), &supplierFilter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, supplierFilter))
	}
}
