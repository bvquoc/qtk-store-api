package ginsupplier

import (
	"net/http"
	"qtk-store-api/common"
	"qtk-store-api/component/appctx"
	"qtk-store-api/component/generator"
	"qtk-store-api/middleware"
	"qtk-store-api/module/supplier/supplierbiz"
	"qtk-store-api/module/supplier/suppliermodel"
	"qtk-store-api/module/supplier/supplierrepo"
	"qtk-store-api/module/supplier/supplierstore"
	"qtk-store-api/module/supplierdebt/supplierdebtstore"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Pay supplier
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path string true "supplier id"
// @Param supplier body suppliermodel.ReqUpdateDebtSupplier true "pay information"
// @Response 200 {object} common.ResSuccess "status of response"
// @Response 400 {object} common.AppError "error"
// @Router /suppliers/{id}/pay [post]
func PaySupplier(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data suppliermodel.ReqUpdateDebtSupplier

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		data.CreateBy = requester.GetUserId()

		db := appCtx.GetMainDBConnection().Begin()

		supplierStore := supplierstore.NewSQLStore(db)
		supplierDebtStore := supplierdebtstore.NewSQLStore(db)
		repo := supplierrepo.NewUpdatePayRepo(supplierStore, supplierDebtStore)

		gen := generator.NewShortIdGenerator()

		business := supplierbiz.NewUpdatePayBiz(gen, repo, requester)

		idSupplierDebt, err := business.PaySupplier(c.Request.Context(), id, &data)

		if err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(idSupplierDebt))
	}
}
