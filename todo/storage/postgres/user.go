package postgres

import (
	"context"

	"github.com/tasnuvatina/grpc-blog/todo/storage"
)

const insertUser =`
	INSERT INTO users(
		user_name,
		email,
		password
	) VALUES(
		:user_name,
		:email,
		:password
	)RETURNING id;
`
func (s *Storage) Register(ctx context.Context, u storage.User) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertUser)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, u); err != nil {
		return 0, err
	}
	return id, nil

}

func (s *Storage) GetUser(ctx context.Context, user_name string) (*storage.User, error) {
	var u storage.User
	if err := s.db.Get(&u, "SELECT * FROM users WHERE user_name=$1", user_name); err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *Storage) GetUserById(ctx context.Context, id int64) (*storage.User, error) {
	var u storage.User
	if err := s.db.Get(&u, "SELECT * FROM users WHERE id=$1", id); err != nil {
		return nil, err
	}
	return &u, nil
}



