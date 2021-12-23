package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
)

func (s *BlogSvc)GetDownvote(ctx context.Context,req *bpb.GetDownvoteRequest) (*bpb.GetDownvoteResponse, error)  {
	id := req.GetBlogID()
	userId := req.GetUserID()

	_,downvoteId,_ := s.core.GetDownvote(context.Background(),id,userId) 
	

	if downvoteId!=0{
		return &bpb.GetDownvoteResponse{
			IsDownvotedId: downvoteId,
		},nil
	}else{
		return &bpb.GetDownvoteResponse{
			IsDownvotedId: 0,
		},nil
	}
}