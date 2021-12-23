package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	
)

func (s *BlogSvc)GetUpvote(ctx context.Context,req *bpb.GetUpvoteRequest) (*bpb.GetUpvoteResponse, error) {
	id := req.GetBlogID()
	userId := req.GetUserID()
	_,upvoteId,_ := s.core.GetUpvote(context.Background(),id,userId)
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