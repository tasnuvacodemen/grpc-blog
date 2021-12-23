package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"github.com/tasnuvatina/grpc-blog/todo/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
func (s *BlogSvc)UpvoteBlog(ctx context.Context,req *bpb.UpvoteBlogRequest) (*bpb.UpvoteBlogResponse, error)  {
	upvote :=storage.Upvote{
		ID: req.GetUpvote().GetID(),
		BlogID: req.GetUpvote().BlogID,
		UserID: req.GetUpvote().UserID,
	}

	id,err := s.core.UpvoteBlog(ctx,upvote)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to Upvote")
	}

	return &bpb.UpvoteBlogResponse{
		UpvoteId: id,
	},nil
}


	
	
	
	
	
	
	
	
	