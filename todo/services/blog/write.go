package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"github.com/tasnuvatina/grpc-blog/todo/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
func (s *BlogSvc)WriteBlog(ctx context.Context,req *bpb.WriteBlogRequest) (*bpb.WriteBlogResponse, error) {
	// need to validate request
	blog:=storage.Blog{
		ID: req.GetBlog().ID,
		AuthorID: req.GetBlog().AuthorID,
		AuthorName: req.GetBlog().AuthorName,
		CreatedAt: req.GetBlog().CreatedAt,
		UpdateAt: req.GetBlog().UpdateAt,
		PictureString: req.GetBlog().PictureString,
		Title: req.GetBlog().Title,
		Description: req.GetBlog().Description,
		UpvoteCount: req.GetBlog().UpvoteCount,
		DownvoteCount: req.GetBlog().DownvoteCount,
		CommentsCount: req.GetBlog().CommentsCount,
	}

	id,err:= s.core.WriteBlog(context.Background(),blog)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to write blog")
	}

	return &bpb.WriteBlogResponse{
		ID: id,
	},nil

}