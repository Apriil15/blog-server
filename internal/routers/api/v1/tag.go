package v1

import (
	"github.com/Apriil15/blog-server/internal/service"
	"github.com/Apriil15/blog-server/pkg/app"
	"github.com/Apriil15/blog-server/pkg/convert"
	"github.com/Apriil15/blog-server/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

// @Summary 新增標籤
// @Produce json
// @Param name body string true "標籤名稱" minlength(3) maxlength(100)
// @Param state body int false "狀態" Enum(0, 1) default(1)
// @Param created_by body string true "建立者" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 刪除標籤
// @Produce json
// @Param id path int true "標籤 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustInt32(),
	}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 更新標籤
// @Produce json
// @Param id path int true "標籤 ID"
// @Param name body string false "標籤名稱" minlength(3) maxlength(100)
// @Param state body int false "狀態" Enum(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{
		ID: convert.StrTo(c.Param("id")).MustInt32(),
	}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (t Tag) Get(c *gin.Context) {}

// @Summary 取得多個標籤
// @Produce json
// @Param name query string false "標籤名稱" maxlength(100)
// @Param state query int false "狀態" Enums(0, 1) default(1)
// @Param page query int false "頁碼"
// @Param page_size query int false "每頁數量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}

	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		//fmt.Printf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	response.ToResponseList(tags, int(totalRows))
}
