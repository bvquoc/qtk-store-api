package categorytransport

import (
	"net/http"
	"qtk-store-api/common"
	"qtk-store-api/component/appctx"
	"qtk-store-api/middleware"
	"qtk-store-api/module/category/categorybiz"
	"qtk-store-api/module/category/categorymodel"
	"qtk-store-api/module/category/categoryrepo"
	"qtk-store-api/module/category/categorystore"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Get all categories
// @Tags categories
// @Accept json
// @Produce json
// @Response 200 {object} categorymodel.ResListCategory
// @Router /categories [get]
func ListCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter categorymodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := categoryrepo.NewListCategoryRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := categorybiz.NewListCategoryRepo(repo, requester)

		result, err := biz.ListCategory(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
