package postgres

import (
	"context"

	"github.com/tasnuvatina/grpc-blog/todo/storage"
)

const insertTodo =`
	INSERT INTO todos(
		title,
		description
	) VALUES(
		:title,
		:description
	)RETURNING id;

`
func (s *Storage) Create(ctx context.Context, t storage.Todo) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertTodo)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, err
	}
	return id, nil

}

func (s *Storage) Get(ctx context.Context, id int64) (*storage.Todo, error) {
	var t storage.Todo
	if err := s.db.Get(&t, "SELECT * FROM todos WHERE id=$1", id); err != nil {
		return nil, err
	}
	return &t, nil

}

const updateTodo = `
	UPDATE todos 
	SET
		title = :title,
		description = :description
	WHERE 
		id = :id
	RETURNING *;
`

func (s *Storage) Update(ctx context.Context, t storage.Todo) (*storage.Todo, error) {
	stmt, err := s.db.PrepareNamed(updateTodo)
	if err != nil {
		return nil, err
	}
	if err := stmt.Get(&t, t); err != nil {
		return nil, err
	}
	return &t, nil

}

func (s *Storage) Delete(ctx context.Context, id int64) error {
	var t storage.Todo
	if err := s.db.Get(&t, "DELETE FROM todos WHERE id=$1 RETURNING *", id); err != nil {
		return err;
	}
	return nil
}