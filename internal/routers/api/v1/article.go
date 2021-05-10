package v1

import (
	"github.com/Apriil15/blog-server/internal/service"
	"github.com/Apriil15/blog-server/pkg/app"
	"github.com/Apriil15/blog-server/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

// @Summary 新增文章
// @Produce json
// @Param title body string true "文章標題" minlength(1) maxlength(100)
// @Param desc body string true "文章簡述" minlength(1) maxlength(100)
// @Param content body string true "文章內容" minlength(1) maxlength(1000)
// @Param cover_image_url body string false "封面圖片網址" minlength(1) maxlength(100)
// @Param state body int false "狀態" Enum(0, 1) default(1)
// @Param created_by body string true "建立者" minlength(3) maxlength(100)
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errorResponse := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errorResponse)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(gin.H{})

}

// @Summary 刪除文章
// @Produce json
// @Param id path int true "文章 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {}

// @Summary 更新文章
// @Produce json
// @Param id path int true "文章 ID"
// @Param title body string false "文章標題" minlength(1) maxlength(100)
// @Param desc body string false "文章簡述" minlength(1) maxlength(100)
// @Param content body string false "文章內容" minlength(1) maxlength(1000)
// @Param cover_image_url body string false "封面圖片網址" minlength(1) maxlength(100)
// @Param state body int false "狀態" Enum(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {}

// @Summary 取得特定文章
// @Produce json
// @Param id path int true "文章 ID"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {}

func (a Article) List(c *gin.Context) {}
