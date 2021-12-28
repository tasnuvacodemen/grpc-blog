package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
) 

func (s *BlogSvc)GetAllComments(ctx context.Context,req *bpb.GetAllCommentsRequest) (*bpb.GetAllCommentsResponse, error)  {
	blogId :=req.GetBlogID()

	res,err := s.core.GetAllComments(context.Background(),blogId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read all the comments for this blog blocks")
	}

	comments := []*bpb.Comment{}

	for _,value :=range res{
		comments = append(comments, &bpb.Comment{
			ID: value.ID,
			BlogID: value.BlogID,
			UserID: value.UserID,
			UserName: value.UserName,
			Content: value.Content,
			CommentedAt: value.CommentedAt,
		})
	}



	return &bpb.GetAllCommentsResponse{
		Comments: comments,
	},nil
}