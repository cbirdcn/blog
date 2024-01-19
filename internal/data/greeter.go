package data

import (
	"blog/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type blogRepo struct {
	data *Data
	log  *log.Helper
}

// NewBlogRepo .
func NewBlogRepo(data *Data, logger log.Logger) biz.BlogRepo {
	return &blogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
