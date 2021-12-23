package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"github.com/tasnuvatina/grpc-blog/todo/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BlogSvc)DownVoteBlog(ctx context.Context,req *bpb.DownVoteRequest) (*bpb.DownVoteResponse, error)  {
	downvote :=storage.Downvote{
		ID: req.GetDownvote().GetID(),
		BlogID: req.GetDownvote().BlogID,
		UserID: req.GetDownvote().UserID,
	}

	id,err := s.core.DownVoteBlog(ctx,downvote)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to downvote")
	}

	return &bpb.DownVoteResponse{
		DownvoteId: id,
	},nil
}