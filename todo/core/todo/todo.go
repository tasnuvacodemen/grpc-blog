package todo

import (
	"context"

	"github.com/tasnuvatina/grpc-blog/todo/storage"
)

type todoStore interface {
	Create(context.Context, storage.Todo) (int64, error)
}

type CoreSvc struct {
	store todoStore
}

func NewCoreSvc(s todoStore) *CoreSvc {
	return &CoreSvc{
		store: s,
	}
}

// our own method

func (cs CoreSvc) Create(ctx context.Context, t storage.Todo) (int64, error) {
	return cs.store.Create(ctx, t)
}
