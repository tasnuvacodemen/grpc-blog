package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
) 

func (s *BlogSvc)ReadAllBlog(ctx context.Context,req *bpb.ReadAllBlogRequest) (*bpb.ReadAllBlogResponse, error) {
	// need to validate request
	
	res,err:= s.core.ReadAllBlog(context.Background())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read all blocks")
	}
	blogs:= []*bpb.Blog{}

	for _, value := range res {
		blogs = append(blogs, &bpb.Blog{
			ID:            value.ID,
			AuthorID:      value.AuthorID,
			AuthorName:    value.AuthorName,
			CreatedAt:     value.CreatedAt,
			UpdateAt:      value.UpdateAt,
			PictureString: value.PictureString,
			Title:         value.Title,
			Description:   value.Description,
			UpvoteCount:   value.UpvoteCount,
			DownvoteCount: value.DownvoteCount,
			CommentsCount: value.CommentsCount,
		
		})
	}
	return &bpb.ReadAllBlogResponse{
		Blogs: blogs,
	},nil
}


	