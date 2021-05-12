package dao

import (
	"time"

	"github.com/Apriil15/blog-server/internal/model"
)

func (d *Dao) CreateArticle(title string, desc string, content string, coverImageUrl string, state uint8, createdBy string) error {
	nowTime := time.Now().Unix()

	article := model.Article{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageUrl,
		State:         state,
		Model: &model.Model{
			CreatedBy:  createdBy,
			CreatedOn:  uint32(nowTime),
			ModifiedOn: uint32(nowTime),
		},
	}
	return article.Create(d.engine)
}

// Update an article
func (d *Dao) UpdateArticle(id uint32, title string, desc string, content string, coverImageUrl string, state uint8, modifiedBy string) error {
	article := model.Article{
		Title:         title,
		Desc:          desc,
		Content:       content,
		State:         state,
		CoverImageUrl: coverImageUrl,
		Model: &model.Model{
			ID:         id,
			ModifiedBy: modifiedBy,
			ModifiedOn: uint32(time.Now().Unix()),
		},
	}
	return article.Update(d.engine)
}

// Delete an article
func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{
		Model: &model.Model{ID: id},
	}
	return article.Delete(d.engine)
}
