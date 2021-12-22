package blog


import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"github.com/tasnuvatina/grpc-blog/todo/storage"
)

type blogCoreStore interface {
	WriteBlog(context.Context,storage.Blog) (int64, error)
	DeleteBlog( context.Context,  int64,  int64)  error
	ReadBlog( context.Context, int64, int64)  (*storage.Blog,bool, error)
	ReadAllBlog( context.Context)  ([]*storage.Blog,error)
	EditBlog( context.Context, storage.Blog)  (*storage.Blog, error)
}
type BlogSvc struct {
	bpb.UnimplementedBlogServiceServer
	core blogCoreStore
}

func NewBlogServer(b blogCoreStore) *BlogSvc {
	return &BlogSvc{
		core: b,
	}
}

