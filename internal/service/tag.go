package service

import (
	"github.com/Apriil15/blog-server/internal/model"
	"github.com/Apriil15/blog-server/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

// Get count of tag
func (s *Service) CountTag(param *CountTagRequest) (int64, error) {
	return s.dao.CountTag(param.Name, param.State)
}

// Get tags
func (s *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return s.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

// Create a tag
func (s *Service) CreateTag(param *CreateTagRequest) error {
	return s.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

// Update a tag
func (s *Service) UpdateTag(param *UpdateTagRequest) error {
	return s.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

// Delete a tag
func (s *Service) DeleteTag(param *DeleteTagRequest) error {
	return s.dao.DeleteTag(param.ID)
}
