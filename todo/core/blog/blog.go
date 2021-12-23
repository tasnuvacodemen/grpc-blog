package blog

import (
	"context"

	"github.com/tasnuvatina/grpc-blog/todo/storage"
)

type blogStore interface {
	WriteBlog(context.Context, storage.Blog) (int64, error)
	DeleteBlog(context.Context, int64, int64) error
	ReadBlog(context.Context, int64, int64) (*storage.Blog, bool, error)
	ReadAllBlog(context.Context) ([]*storage.Blog, error)
	EditBlog(context.Context, storage.Blog) (*storage.Blog, error)

	UpvoteBlog(context.Context, storage.Upvote) (int64, error)
	RevertUpvoteBlog(context.Context, int64, int64) error
	GetUpvote(context.Context, int64, int64) (*storage.Upvote, int64, error)
	GetAllUpvote(context.Context, int64) ([]*storage.Upvote, error)
	DownVoteBlog(context.Context, storage.Downvote) (int64, error)
	RevertDownVoteBlog(context.Context, int64, int64) error
	GetDownvote(context.Context, int64, int64) (*storage.Downvote, int64, error)
	GetAllDownvote(context.Context, int64) ([]*storage.Downvote, error)
	CommentBlog(context.Context, storage.Comment) (int64, error)
	GetAllComments(context.Context, int64) ([]*storage.Comment, error)
}

type BlogCoreSvc struct {
	store blogStore
}

func NewBlogCoreSvc(b blogStore) *BlogCoreSvc {
	return &BlogCoreSvc{
		store: b,
	}
}

// our own method
func (bs BlogCoreSvc) WriteBlog(ctx context.Context, b storage.Blog) (int64, error) {
	return bs.store.WriteBlog(ctx, b)
}
func (bs BlogCoreSvc) DeleteBlog(ctx context.Context, id int64, author_id int64) error {
	return bs.store.DeleteBlog(ctx, id, author_id)
}
func (bs BlogCoreSvc) ReadBlog(ctx context.Context, id int64, author_id int64) (*storage.Blog, bool, error) {
	return bs.store.ReadBlog(ctx, id, author_id)
}
func (bs BlogCoreSvc) ReadAllBlog(ctx context.Context) ([]*storage.Blog, error) {
	return bs.store.ReadAllBlog(ctx)
}
func (bs BlogCoreSvc) EditBlog(ctx context.Context, b storage.Blog) (*storage.Blog, error) {
	return bs.store.EditBlog(ctx, b)
}




func (bs BlogCoreSvc) UpvoteBlog(ctx context.Context, u storage.Upvote) (int64, error) {
	return bs.store.UpvoteBlog(ctx, u)
}
func (bs BlogCoreSvc) RevertUpvoteBlog(ctx context.Context, upvote_id int64, user_id int64) error {
	return bs.store.RevertUpvoteBlog(ctx, upvote_id, user_id)
}
func (bs BlogCoreSvc) GetUpvote(ctx context.Context, blog_id int64, user_id int64) (*storage.Upvote, int64, error) {
	return bs.store.GetUpvote(ctx, blog_id, user_id)
}
func (bs BlogCoreSvc) GetAllUpvote(ctx context.Context, blog_id int64) ([]*storage.Upvote, error) {
	return bs.store.GetAllUpvote(ctx, blog_id)
}
func (bs BlogCoreSvc) DownVoteBlog(ctx context.Context, u storage.Downvote) (int64, error) {
	return bs.store.DownVoteBlog(ctx, u)
}
func (bs BlogCoreSvc) RevertDownVoteBlog(ctx context.Context, downvote_id int64, user_id int64) error {
	return bs.store.RevertDownVoteBlog(ctx, downvote_id, user_id)
}
func (bs BlogCoreSvc) GetDownvote(ctx context.Context, blog_id int64, user_id int64) (*storage.Downvote, int64, error) {
	return bs.store.GetDownvote(ctx, blog_id, user_id)
}
func (bs BlogCoreSvc) GetAllDownvote(ctx context.Context, blog_id int64) ([]*storage.Downvote, error) {
	return bs.store.GetAllDownvote(ctx, blog_id)
}
func (bs BlogCoreSvc) CommentBlog(ctx context.Context, u storage.Comment) (int64, error) {
	return bs.store.CommentBlog(ctx, u)
}
func (bs BlogCoreSvc) GetAllComments(ctx context.Context, blog_id int64) ([]*storage.Comment, error) {
	return bs.store.GetAllComments(ctx, blog_id)
}
