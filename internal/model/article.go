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

// Get articles
func (a *Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize) // 跳過幾個，然後取幾個
	}

	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}

	db = db.Where("state = ?", a.State)
	if err = db.Where("is_del = ?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

// Get article count
func (a *Article) Count(db *gorm.DB) (int64, error) {
	var count int64
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}

	db = db.Where("state = ?", a.State)
	err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
