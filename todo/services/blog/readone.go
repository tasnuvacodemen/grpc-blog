package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BlogSvc) ReadBlog(ctx context.Context, req *bpb.ReadBlogRequest) (*bpb.ReadBlogResponse, error) {
	// need to validate request
	blogId := req.GetBlogID()
	userId := req.GetAuthorID()
	blog, isAuthor, err := s.core.ReadBlog(context.Background(), blogId, userId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to read blog")
	}
	return &bpb.ReadBlogResponse{
		Blog: &bpb.Blog{
			ID:            blog.ID,
			AuthorID:      blog.AuthorID,
			AuthorName:    blog.AuthorName,
			CreatedAt:     blog.CreatedAt,
			UpdateAt:      blog.UpdateAt,
			PictureString: blog.PictureString,
			Title:         blog.Title,
			Description:   blog.Description,
			UpvoteCount:   blog.UpvoteCount,
			DownvoteCount: blog.DownvoteCount,
			CommentsCount: blog.CommentsCount,
		},
		IsAuthor: isAuthor,
		
	},nil
}
