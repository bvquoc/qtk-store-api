package producttransport

import (
	"qtk-store-api/component/appctx"

	"github.com/gin-gonic/gin"
)

func CreateProduct(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//var data productmodel.ReqCreateProduct
		////c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
		//
		//if err := c.ShouldBind(&data); err != nil {
		//	panic(common.ErrInvalidRequest(err))
		//}
		//
		//requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)
		//
		//db := appCtx.GetMainDBConnection().Begin()
		//
		//store := productstore.NewSQLStore(db)
		//repo := productrepo.NewCreateProductRepo(store)
		//
		//gen := generator.NewShortIdGenerator()
		//
		//business := productbiz.NewCreateProductBiz(gen, repo, requester)
		//
		//fmt.Print(data)
		//if err := business.CreateProduct(c.Request.Context(), &data); err != nil {
		//	db.Rollback()
		//	panic(err)
		//}
		//
		//if err := db.Commit().Error; err != nil {
		//	db.Rollback()
		//	panic(err)
		//}
		//
		//c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.ID))
	}
}
