package middleware

import (
	"errors"
	"fmt"
	"qtk-store-api/common"
	"qtk-store-api/component/appctx"
	"qtk-store-api/component/tokenprovider/jwt"
	"qtk-store-api/module/role/rolemodel"
	"qtk-store-api/module/user/userstore"
	"strings"

	"github.com/gin-gonic/gin"
)

type Requester interface {
	GetUserId() string
	GetEmail() string
	GetRole() rolemodel.Role
	IsHasFeature(featureCode string) bool
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//Authorization : Bearn{token}
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

func RequireAuth(appCtx appctx.AppContext) func(ctx *gin.Context) {

	tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()

		store := userstore.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(
			c.Request.Context(),
			map[string]interface{}{
				"id": payload.UserId,
			},
			"Role.RoleFeatures",
		)

		if err != nil {
			//c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			panic(err)
		}

		if !user.IsActive {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		c.Set(common.CurrentUserStr, user)
		c.Next()
	}

}
