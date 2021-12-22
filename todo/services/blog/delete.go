package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
) 

func (s *BlogSvc)DeleteBlog(ctx context.Context,req *bpb.DeleteBlogRequest) (*bpb.DeleteBlogResponse, error) {
	// need to validate request
	id:=req.GetID()
	user_id :=req.GetAuthorID()
	if err := s.core.DeleteBlog(context.Background(),id,user_id);err!=nil{
		return nil,status.Errorf(codes.Internal, "Failed to delete blog")
	}
	return &bpb.DeleteBlogResponse{},nil
}

