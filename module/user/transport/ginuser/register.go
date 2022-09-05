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

func UserRegister(appctx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user usermodel.UserCreate
		if err := ctx.ShouldBind(&user); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appctx.GetMysqlConnection()
		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()
		bus := userbusiness.NewRegisterUserBusiness(store, md5)
		if err := bus.Register(ctx.Request.Context(), &user); err != nil {
			panic(err)
		}
		user.GenUID(2)
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(user.FakeID.String()))
	}
}
