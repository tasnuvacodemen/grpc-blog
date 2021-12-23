package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
) 

func (s *BlogSvc)GetAllUpvote(ctx context.Context,req *bpb.GetAllUpvoteRequest) (*bpb.GetAllUpvoteResponse, error)  {
	blogId := req.GetBlogID()

	res,err := s.core.GetAllUpvote(context.Background(),blogId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read all the upvotes for this blog blocks")
	}

	upvotes :=[]*bpb.Upvote{}

	for _,value := range res {
		upvotes = append(upvotes, &bpb.Upvote{
			ID: value.ID,
			BlogID: value.BlogID,
			UserID: value.UserID,
		})
	}

	return &bpb.GetAllUpvoteResponse{
		Upvotes: upvotes,
	},nil
}