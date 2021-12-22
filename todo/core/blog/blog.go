package blog

import (
	"context"

	"github.com/tasnuvatina/grpc-blog/todo/storage"
)

type blogStore interface {
	WriteBlog(context.Context,storage.Blog) (int64, error)
	DeleteBlog( context.Context,  int64,  int64)  error
	ReadBlog( context.Context, int64, int64)  (*storage.Blog,bool, error)
	ReadAllBlog( context.Context)  ([]*storage.Blog,error)
	EditBlog( context.Context, storage.Blog)  (*storage.Blog, error)
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
func (bs BlogCoreSvc) DeleteBlog(ctx context.Context, id int64, author_id int64)  error {
	return bs.store.DeleteBlog(ctx, id,author_id)
}
func (bs BlogCoreSvc) ReadBlog(ctx context.Context, id int64, author_id int64) (*storage.Blog,bool, error) {
	return bs.store.ReadBlog(ctx, id,author_id)
}
func (bs BlogCoreSvc) ReadAllBlog(ctx context.Context) ([]*storage.Blog,error) {
	return bs.store.ReadAllBlog(ctx)
}
func (bs BlogCoreSvc) EditBlog(ctx context.Context, b storage.Blog) (*storage.Blog, error) {
	return bs.store.EditBlog(ctx, b)
}
