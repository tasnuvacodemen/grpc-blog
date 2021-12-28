package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
) 

func (s *BlogSvc)GetAllCommentCount(ctx context.Context,req *bpb.GetAllCommentCountRequest) (*bpb.GetAllCommentCountResponse, error)  {
	blogId := req.GetBlogID()

	commentCount,err :=s.core.GetAllCommentCount(context.Background(),blogId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read all comments count for this blog blocks")
	}
																									

	return &bpb.GetAllCommentCountResponse{
		CommentCount: commentCount,
	},nil
}