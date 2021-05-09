package v1

import (
	"github.com/Apriil15/blog-server/internal/service"
	"github.com/Apriil15/blog-server/pkg/app"
	"github.com/Apriil15/blog-server/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// Get JWT token
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedTokenGenetate)
		return
	}

	response.ToResponse(gin.H{"token": token})
}
