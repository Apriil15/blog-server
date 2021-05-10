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
