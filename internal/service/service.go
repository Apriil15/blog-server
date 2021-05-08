package service

import (
	"context"

	"github.com/Apriil15/blog-server/global"
	"github.com/Apriil15/blog-server/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	service := Service{ctx: ctx}
	service.dao = dao.New(global.DBEngine)
	return service
}
