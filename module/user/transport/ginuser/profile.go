package ginuser

import (
	"food/common"
	"food/component/appctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(appctx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := ctx.MustGet("user")
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
