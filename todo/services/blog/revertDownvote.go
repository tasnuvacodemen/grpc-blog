package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BlogSvc)RevertDownVoteBlog(ctx context.Context,req *bpb.RevertDownVoteBlogRequest) (*bpb.RevertDownVoteBlogResponse, error) {
	id:=req.GetDownvoteId()
	user_id :=req.GetUserId()
	if err := s.core.RevertDownVoteBlog(context.Background(),id,user_id);err!=nil{
		return nil,status.Errorf(codes.Internal, "Failed to revert downvote")
	}
	return &bpb.RevertDownVoteBlogResponse{},nil
}