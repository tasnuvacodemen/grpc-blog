package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
) 

func (s *BlogSvc)GetAllUpvoteCount(ctx context.Context,req *bpb.GetAllUpvoteCountRequest) (*bpb.GetAllUpvoteCountResponse, error)  {
	blogId := req.GetBlogID()

	upvoteCount,err :=s.core.GetAllUpvoteCount(context.Background(),blogId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read all upvote count for this blog blocks")
	}
																									

	return &bpb.GetAllUpvoteCountResponse{
		UpvoteCount: upvoteCount,
	},nil
}