package routers

import (
	_ "github.com/Apriil15/blog-server/docs"
	"github.com/Apriil15/blog-server/internal/middleware"
	"github.com/Apriil15/blog-server/internal/routers/api"
	v1 "github.com/Apriil15/blog-server/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()

	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", v1.GetAuth)
	r.POST("/upload/file", upload.UploadFile)

	api := r.Group("api/v1")
	{
		api.Use(middleware.JWT())

		api.POST("tags", tag.Create)
		api.DELETE("tags/:id", tag.Delete)
		api.PUT("tags/:id", tag.Update)
		api.PATCH("tags/:id/state", tag.Update)
		api.GET("tags", tag.List)

		api.POST("articles", article.Create)
		api.DELETE("articles/:id", article.Delete)
		api.PUT("articles/:id", article.Update)
		api.PATCH("articles/:id/state", article.Update)
		api.GET("articles/:id", article.Get)
		api.GET("articles", article.List)
	}

	return r
}
