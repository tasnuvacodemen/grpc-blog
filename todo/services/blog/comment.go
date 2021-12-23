package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"github.com/tasnuvatina/grpc-blog/todo/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BlogSvc)CommentBlog(ctx context.Context,req *bpb.CommentBlogRequest) (*bpb.CommentBlogResponse, error)  {
	comment :=storage.Comment{
		ID: req.GetComment().ID,
		BlogID: req.GetComment().BlogID,
		UserID: req.GetComment().UserID,
		UserName: req.GetComment().UserName,
		Content: req.GetComment().Content,
		CommentedAt: req.GetComment().CommentedAt,
	}

	id,err := s.core.CommentBlog(ctx,comment)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to Upvote")
	}

	return &bpb.CommentBlogResponse{
		CommentID: id,
	},nil
}