package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type Blog struct {
}

type BlogRepo interface {
}

type BlogUsecase struct {
	repo BlogRepo
	log  *log.Helper
}

func NewBlogUsecase(repo BlogRepo, logger log.Logger) *BlogUsecase {
	return &BlogUsecase{repo: repo, log: log.NewHelper(logger)}
}
