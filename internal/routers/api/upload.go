package api

import (
	"github.com/Apriil15/blog-server/internal/service"
	"github.com/Apriil15/blog-server/pkg/app"
	"github.com/Apriil15/blog-server/pkg/convert"
	"github.com/Apriil15/blog-server/pkg/errcode"
	"github.com/Apriil15/blog-server/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

// Upload file
func (u *Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()

	if err != nil {
		errResponse := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errResponse)
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		errResponse := errcode.ErrorUploadFileUrl.WithDetails(err.Error())
		response.ToErrorResponse(errResponse)
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
