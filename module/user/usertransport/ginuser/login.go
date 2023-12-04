package ginuser

import (
	"net/http"
	"qtk-store-api/common"
	"qtk-store-api/component/appctx"
	"qtk-store-api/component/hasher"
	"qtk-store-api/component/tokenprovider/jwt"
	"qtk-store-api/module/user/userbiz"
	"qtk-store-api/module/user/usermodel"
	"qtk-store-api/module/user/userrepo"
	"qtk-store-api/module/user/userstore"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Security BearerAuth
// @Summary Login
// @Tags common
// @Accept json
// @Produce json
// @Param user body usermodel.ReqLoginUser true "login information"
// @Response 200 {object} usermodel.Account "user token"
// @Response 400 {object} common.AppError "error"
// @Router /login [post]
func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.ReqLoginUser

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection().Begin()

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

		store := userstore.NewSQLStore(db)
		repo := userrepo.NewLoginRepo(store)

		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBiz(appCtx, repo, 60*60*24*30, tokenProvider, md5)
		account, err := business.Login(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
