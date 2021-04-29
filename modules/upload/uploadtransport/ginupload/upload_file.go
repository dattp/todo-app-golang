package ginupload

import (
	"github.com/gin-gonic/gin"
	_ "image/jpeg"
	_ "image/png"
	"todo-app/common"
	"todo-app/component"
	"todo-app/modules/upload/uploadbiz"
)

func Upload(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		//db := ctx.GetMainDBConnection();
		fileHeader, err := c.FormFile("file")
		//fileHeader
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		//c.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename)) // luu local

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)

		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(img))

		//c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
