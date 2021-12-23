package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BlogSvc)GetDownvote(ctx context.Context,req *bpb.GetDownvoteRequest) (*bpb.GetDownvoteResponse, error)  {
	id := req.GetBlogID()
	userId := req.GetUserID()

	_,downvoteId,err := s.core.GetDownvote(context.Background(),id,userId) 
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to your read downvote foe this blog")
	}

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