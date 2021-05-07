package model

import "github.com/Apriil15/blog-server/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// TagSwagger is a struct for swagger page.
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t *Tag) TableName() string {
	return "blog_tag"
}
