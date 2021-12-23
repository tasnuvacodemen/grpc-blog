package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
) 

func (s *BlogSvc)RevertUpvoteBlog(ctx context.Context,req *bpb.RevertUpvoteBlogRequest) (*bpb.RevertUpvoteBlogResponse, error)  {
	id :=req.GetUpvoteId()
	user_id :=req.GetUserId()

	if err := s.core.RevertUpvoteBlog(context.Background(),id,user_id);err!=nil{
		return nil,status.Errorf(codes.Internal, "Failed to revert upvote")
	}
	return &bpb.RevertUpvoteBlogResponse{},nil
}