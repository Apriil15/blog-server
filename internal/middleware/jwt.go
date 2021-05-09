package middleware

import (
	"strings"

	"github.com/Apriil15/blog-server/pkg/app"
	"github.com/Apriil15/blog-server/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is a HandlerFunc for middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		var code = errcode.Success

		// get token from query / header / Authorization
		if s, exist := c.GetQuery("token"); exist {
			token = s
		}
		if s := c.GetHeader("token"); s != "" {
			token = s
		}
		if s := c.Request.Header.Get("Authorization"); s != "" {
			splitToken := strings.Split(s, "Bearer ")
			token = splitToken[1]
		}

		// check token
		if token == "" {
			code = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = errcode.UnauthorizedTokenTimeout
				default:
					code = errcode.UnauthorizedTokenError
				}
			}
		}

		if code != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(code)
			c.Abort()
			return
		}

		c.Next()
	}
}
