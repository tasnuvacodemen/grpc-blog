package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
) 

func (s *BlogSvc)ReadAllSearchedBlog(ctx context.Context,req *bpb.ReadAllBlogSearchedRequest) (*bpb.ReadAllBlogSearchedResponse, error) {
	// need to validate request
	searchValue :=req.GetSearchValue()
	res,err:= s.core.ReadAllSearchedBlog(context.Background(),searchValue)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read all searched blogs")
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
	return &bpb.ReadAllBlogSearchedResponse{
		Blogs: blogs,
	},nil
}