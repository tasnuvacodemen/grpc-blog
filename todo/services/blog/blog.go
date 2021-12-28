package blog

import (
	"context"

	bpb "github.com/tasnuvatina/grpc-blog/proto/blog"
	"github.com/tasnuvatina/grpc-blog/todo/storage"
)

type blogCoreStore interface {
	WriteBlog(context.Context, storage.Blog) (int64, error)
	DeleteBlog(context.Context, int64, int64) error
	ReadBlog(context.Context, int64, int64) (*storage.Blog, bool, error)
	ReadAllBlog(context.Context) ([]*storage.Blog, error)
	ReadAllSearchedBlog(context.Context,string) ([]*storage.Blog, error)
	EditBlog(context.Context, storage.Blog) (*storage.Blog, error)

	UpvoteBlog(context.Context, storage.Upvote) (int64, error)
	RevertUpvoteBlog(context.Context, int64, int64) error
	GetUpvote(context.Context, int64, int64) (*storage.Upvote, int64, error)
	GetAllUpvote(context.Context, int64) ([]*storage.Upvote, error)
	GetAllUpvoteCount(context.Context, int64) (int64, error)
	DownVoteBlog(context.Context, storage.Downvote) (int64, error)
	RevertDownVoteBlog(context.Context, int64, int64) error
	GetDownvote(context.Context, int64, int64) (*storage.Downvote, int64, error)
	GetAllDownvote(context.Context, int64) ([]*storage.Downvote, error)
	GetAllDownvoteCount(context.Context, int64) (int64, error)
	CommentBlog(context.Context, storage.Comment) (int64, error)
	GetAllComments(context.Context, int64) ([]*storage.Comment, error)
	GetAllCommentCount(context.Context, int64) (int64, error)
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
