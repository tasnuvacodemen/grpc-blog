package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
) 

func (s *BlogSvc)GetAllDownvote(ctx context.Context,req *bpb.GetAllDownvoteRequest) (*bpb.GetAllDownvoteResponse, error){
	blogId := req.GetBlogID()

	res,err := s.core.GetAllDownvote(context.Background(),blogId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read all the downvotes for this blog blocks")
	}

	downvotes :=[]*bpb.Downvote{}

	for _,value := range res {
		downvotes = append(downvotes, &bpb.Downvote{
			ID: value.ID,
			BlogID: value.BlogID,
			UserID: value.UserID,
		})
	}

	return &bpb.GetAllDownvoteResponse{
		Downvotes: downvotes,
	},nil
}