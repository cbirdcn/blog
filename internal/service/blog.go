package service

import (
	pb "blog/api/blog/v1"
	"blog/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type BlogService struct {
	pb.UnimplementedBlogServer
	uc  *biz.BlogUsecase
	log *log.Helper
}

func NewBlogService(uc *biz.BlogUsecase, logger log.Logger) *BlogService {
	return &BlogService{uc: uc, log: log.NewHelper(logger)}
}

func (s *BlogService) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogReply, error) {
	return &pb.CreateBlogReply{}, nil
}
func (s *BlogService) UpdateBlog(ctx context.Context, req *pb.UpdateBlogRequest) (*pb.UpdateBlogReply, error) {
	return &pb.UpdateBlogReply{}, nil
}
func (s *BlogService) DeleteBlog(ctx context.Context, req *pb.DeleteBlogRequest) (*pb.DeleteBlogReply, error) {
	return &pb.DeleteBlogReply{}, nil
}
func (s *BlogService) GetBlog(ctx context.Context, req *pb.GetBlogRequest) (*pb.GetBlogReply, error) {
	return &pb.GetBlogReply{}, nil
}
func (s *BlogService) ListBlog(ctx context.Context, req *pb.ListBlogRequest) (*pb.ListBlogReply, error) {
	s.log.WithContext(ctx).Infof("CreateBlog ...")
	return &pb.ListBlogReply{}, nil
}