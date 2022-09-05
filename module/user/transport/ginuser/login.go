package ginuser

import (
	"food/common"
	"food/component/appctx"
	"food/component/hasher"
	userbusiness "food/module/user/business"
	usermodel "food/module/user/model"
	userstorage "food/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appctx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user usermodel.UserLogin
		if err := ctx.ShouldBind(&user); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := userstorage.NewSqlStore(appctx.GetMysqlConnection())
		hasher := hasher.NewMd5Hash()
		bus := userbusiness.NewLoginBusiness(store, appctx.GetTokenProvider(), hasher, 60*60*24)
		token, err := bus.Login(ctx.Request.Context(), &user)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(token))
	}
}
