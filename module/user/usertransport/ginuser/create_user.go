package ginuser

import (
	"net/http"
	"qtk-store-api/common"
	"qtk-store-api/component/appctx"
	"qtk-store-api/component/generator"
	"qtk-store-api/component/hasher"
	"qtk-store-api/middleware"
	"qtk-store-api/module/role/rolestore"
	"qtk-store-api/module/user/userbiz"
	"qtk-store-api/module/user/usermodel"
	"qtk-store-api/module/user/userrepo"
	"qtk-store-api/module/user/userstore"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body usermodel.ReqCreateUser true "user need to create"
// @Response 200 {object} usermodel.ResCreateUser "user id"
// @Response 400 {object} common.AppError "error"
// @Router /users [post]
func CreateUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.ReqCreateUser

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		userStore := userstore.NewSQLStore(db)
		roleStore := rolestore.NewSQLStore(db)
		repo := userrepo.NewCreateUserRepo(userStore, roleStore)

		md5 := hasher.NewMd5Hash()
		gen := generator.NewShortIdGenerator()
		biz := userbiz.NewCreateUserBiz(gen, repo, md5, requester)

		if err := biz.CreateUser(c.Request.Context(), &data); err != nil {
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
