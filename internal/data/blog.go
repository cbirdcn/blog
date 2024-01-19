package data

import (
	"context"
	"blog/internal/biz"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

// 1、作为biz中repo接口的实现
// 2、调用data中的db操作

type blogRepo struct {
	data *Data
	log  *log.Helper
}

func (b blogRepo) CreateBlog(ctx context.Context, blog *biz.Blog) error {
	panic("implement me")
}

func (b blogRepo) DeleteBlog(ctx context.Context, id int64) error {
	panic("implement me")
}

func (b blogRepo) UpdateBlog(ctx context.Context, id int64, blog *biz.Blog) error {
	panic("implement me")
}

func (b blogRepo) GetBlog(ctx context.Context, id int64) (*biz.Blog, error) {

	get := b.data.rdb.HGetAll(ctx, "blog:" + strconv.FormatInt(id, 10))
	rv, _ := get.Result()
	fmt.Println(rv)

	ID, _ := strconv.ParseInt(rv["ID"], 10, 64)
	CreatedAt, _ := time.ParseInLocation("2006-01-02", rv["CreatedAt"], time.Local)
	UpdatedAt, _ := time.ParseInLocation("2006-01-02", rv["UpdatedAt"], time.Local)
	Like, _ := b.GetBlogLike(ctx, id)
	return &biz.Blog{
		ID:        ID,
		Title:     rv["Title"],
		Content:   rv["Content"],
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
		Like: Like,
	}, nil

	//if err == redis.Nil {
	//}
	//if err != nil {
	//	return nil, err
	//}
	//
	////panic("implement me")
	//return *biz.Blog{
	//	ID:        0,
	//	Title:     "1",
	//	Content:   "2",
	//	CreatedAt: time.Now(),
	//	UpdatedAt: time.Now(),
	//	Like:      0,
	//}, nil
}

func (b blogRepo) ListBlog(ctx context.Context) ([]*biz.Blog, error) {
	panic("implement me")
}

//func (b blogRepo) GetBlogLike(ctx context.Context, id int64) (rv int64, err error) {
//	panic("implement me")
//}
//
//func (b blogRepo) IncBlogLike(ctx context.Context, id int64) error {
//	panic("implement me")
//}

// NewBlogRepo .
func NewBlogRepo(data *Data, logger log.Logger) biz.BlogRepo {
	return &blogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func likeKey(id int64) string {
	return fmt.Sprintf("like:%d", id)
}

func (b *blogRepo) GetBlogLike(ctx context.Context, id int64) (rv int64, err error) {
	get := b.data.rdb.Get(ctx, likeKey(id))
	rv, err = get.Int64()
	fmt.Println(rv)
	if err == redis.Nil {
		return 0, nil
	}
	return
}

func (b *blogRepo) IncBlogLike(ctx context.Context, id int64) error {
	_, err := b.data.rdb.Incr(ctx, likeKey(id)).Result()
	return err
}
