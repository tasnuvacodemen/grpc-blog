package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
) 

func (s *BlogSvc)GetAllDownvoteCount(ctx context.Context,req *bpb.GetAllDownvoteCountRequest) (*bpb.GetAllDownvoteCountResponse, error)  {
	blogId := req.GetBlogID()

	downvoteCount,err :=s.core.GetAllDownvoteCount(context.Background(),blogId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read all downvote count for this blog blocks")
	}
																									

	return &bpb.GetAllDownvoteCountResponse{
		DownvoteCount: downvoteCount,
	},nil
}