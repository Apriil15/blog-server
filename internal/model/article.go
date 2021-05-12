package model

import (
	"github.com/Apriil15/blog-server/pkg/app"
	"gorm.io/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

// ArticleSwagger is a struct for swagger page.
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a *Article) TableName() string {
	return "blog_article"
}

func (a *Article) Create(db *gorm.DB) error {

	return db.Create(&a).Error
}

// Delete an article
func (a *Article) Delete(db *gorm.DB) error {
	return db.Where("is_del = ?", 0).Delete(&a).Error
}

// Update an article
func (a *Article) Update(db *gorm.DB) error {
	values := map[string]interface{}{
		"state":       a.State,
		"modified_by": a.ModifiedBy,
		"modified_on": a.ModifiedOn,
	}
	if a.Title != "" {
		values["title"] = a.Title
	}
	if a.Desc != "" {
		values["desc"] = a.Desc
	}
	if a.Content != "" {
		values["content"] = a.Content
	}
	if a.CoverImageUrl != "" {
		values["cover_image_url"] = a.CoverImageUrl
	}

	return db.Model(&a).Where("is_del = ?", 0).Updates(values).Error
}
