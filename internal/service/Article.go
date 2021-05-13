package service

import (
	"github.com/Apriil15/blog-server/internal/model"
	"github.com/Apriil15/blog-server/pkg/app"
)

type CountArticleRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=1,max=100"`
	Desc          string `form:"desc" binding:"required,min=1,max=100"`
	Content       string `form:"content" binding:"required,min=1,max=1000"`
	CoverImageUrl string `form:"cover_image_url" binding:"min=0,max=100"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"min=0,max=100"`
	Desc          string `form:"desc" binding:"min=0,max=100"`
	State         uint8  `form:"state" binding:"required,oneof=0 1"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=1,max=100"`
	Content       string `form:"content" binding:"min=0,max=1000"`
	CoverImageUrl string `form:"cover_image_url" binding:"min=0,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

// Create an article
func (s *Service) CreateArticle(param *CreateArticleRequest) error {
	return s.dao.CreateArticle(param.Title, param.Desc, param.Content, param.CoverImageUrl, param.State, param.CreatedBy)
}

// Update an article
func (s *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return s.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.Content, param.CoverImageUrl, param.State, param.ModifiedBy)
}

// Delete an article
func (s *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return s.dao.DeleteArticle(param.ID)
}

// Get articles
func (s *Service) GetArticles(param *ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return s.dao.GetArticles(param.Title, param.State, pager.Page, pager.PageSize)
}

// Get article count
func (s *Service) CountArticle(param *CountArticleRequest) (int64, error) {
	return s.dao.CountArticle(param.Title, param.State)
}
