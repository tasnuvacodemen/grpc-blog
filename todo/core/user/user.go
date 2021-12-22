package user

import (
	"context"

	"github.com/tasnuvatina/grpc-blog/todo/storage"
)

type userStore interface {
	Register(context.Context, storage.User) (int64, error)
	GetUser(context.Context, string) (*storage.User, error)
	GetUserById(context.Context, int64) (*storage.User, error)
}

type UserCoreSvc struct {
	store userStore
}

func NewUserCoreSvc(s userStore) *UserCoreSvc {
	return &UserCoreSvc{
		store: s,
	}
}

// our own method

func (cs UserCoreSvc) Register(ctx context.Context, u storage.User) (int64, error) {
	return cs.store.Register(ctx, u)
}

func (cs UserCoreSvc) GetUser(ctx context.Context, user_name string) (*storage.User, error) {
	return cs.store.GetUser(ctx, user_name)
}

func (cs UserCoreSvc) GetUserById(ctx context.Context,id int64) (*storage.User, error) {
	return cs.store.GetUserById(ctx, id)
}


