package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BlogSvc)GetUpvote(ctx context.Context,req *bpb.GetUpvoteRequest) (*bpb.GetUpvoteResponse, error) {
	id := req.GetBlogID()
	userId := req.GetUserID()

	_,upvoteId,err := s.core.GetUpvote(context.Background(),id,userId) 
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read your upvotes foe this blog")
	}

	if upvoteId!=0{
		return &bpb.GetUpvoteResponse{
			IsUpvotedId: upvoteId,
		},nil
	}else{
		return &bpb.GetUpvoteResponse{
			IsUpvotedId: 0,
		},nil
	}

	
}