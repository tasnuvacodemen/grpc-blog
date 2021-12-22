package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"github.com/tasnuvatina/grpc-blog/todo/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BlogSvc) EditBlog(ctx context.Context, req *bpb.EditBlogRequest) (*bpb.EditBlogResponse, error) {
	// need to validate request
	blog := storage.Blog{
		ID:            req.GetBlog().ID,
		AuthorID:      req.GetBlog().AuthorID,
		AuthorName:    req.GetBlog().AuthorName,
		CreatedAt:     req.GetBlog().CreatedAt,
		UpdateAt:      req.GetBlog().UpdateAt,
		PictureString: req.GetBlog().PictureString,
		Title:         req.GetBlog().Title,
		Description:   req.GetBlog().Description,
		UpvoteCount:   req.GetBlog().UpvoteCount,
		DownvoteCount: req.GetBlog().DownvoteCount,
		CommentsCount: req.GetBlog().CommentsCount,
	}
	editedBlog,err:= s.core.EditBlog(context.Background(),blog)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to edit blog")
	}
	return &bpb.EditBlogResponse{
		Blog: &bpb.Blog{
		ID:            editedBlog.ID,
		AuthorID:      editedBlog.AuthorID,
		AuthorName:    editedBlog.AuthorName,
		CreatedAt:     editedBlog.CreatedAt,
		UpdateAt:      editedBlog.UpdateAt,
		PictureString: editedBlog.PictureString,
		Title:         editedBlog.Title,
		Description:   editedBlog.Description,
		UpvoteCount:   editedBlog.UpvoteCount,
		DownvoteCount: editedBlog.DownvoteCount,
		CommentsCount: editedBlog.CommentsCount,
		},
	},nil
}
