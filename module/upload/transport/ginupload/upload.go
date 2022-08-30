package ginupload

import (
	"food/common"
	"food/component/appctx"
	uploadbusiness "food/module/upload/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(app appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrCannotReadFile(err))
		}
		defer file.Close()

		data := make([]byte, fileHeader.Size)

		if _, err := file.Read(data); err != nil {
			panic(common.ErrCannotReadFile(err))
		}
		folder := ctx.DefaultPostForm("folder", "img")
		bus := uploadbusiness.NewUploadBus(app.GetUploadProvider())
		img, err := bus.Upload(ctx.Request.Context(), data, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
