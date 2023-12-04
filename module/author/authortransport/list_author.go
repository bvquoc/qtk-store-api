package authortransport

import (
	"net/http"
	"qtk-store-api/common"
	"qtk-store-api/component/appctx"
	"qtk-store-api/middleware"
	"qtk-store-api/module/author/authorbiz"
	"qtk-store-api/module/author/authormodel"
	"qtk-store-api/module/author/authorrepo"
	"qtk-store-api/module/author/authorstore"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Get all authors
// @Tags authors
// @Accept json
// @Produce json
// @Response 200 {object} authormodel.ResListAuthor
// @Router /authors [get]
// @Param page query int false "Page"
// @Param limit query int false "Limit"
func ListAuthor(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter authormodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := authorstore.NewSQLStore(appCtx.GetMainDBConnection())
		repo := authorrepo.NewListAuthorRepo(store)

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		biz := authorbiz.NewListAuthorRepo(repo, requester)

		result, err := biz.ListAuthor(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
