package middleware

import (
	"errors"
	"food/common"
	"food/component/appctx"
	userstorage "food/module/user/storage"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(appctx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			panic(common.ErrJWT(nil))
		}
		t := strings.Split(token, " ")
		payload, err := appctx.GetTokenProvider().Validate(t[1])
		if err != nil {
			panic(common.ErrJWT(err))
		}
		db := userstorage.NewSqlStore(appctx.GetMysqlConnection())
		user, err := db.FindUser(ctx.Request.Context(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			panic(err)
		}
		if user.Status == 0 {
			panic(common.ErrLoginFail(errors.New("account was banned")))
		}
		user.GenUID(2)
		ctx.Set("user", user)
		ctx.Next()
	}
}
