package postgres

import (
	"context"

	"github.com/tasnuvatina/grpc-blog/todo/storage"
)
const writeBlog =`
	INSERT INTO blogs(
		author_id,
		author_name,
		created_at,
		updated_at,
		picture_string,
		title,
		description,
		upvote_count,
		downvote_count,
		comment_count
	) VALUES(
		:author_id,
		:author_name,
		:created_at,
		:updated_at,
		:picture_string,
		:title,
		:description,
		:upvote_count,
		:downvote_count,
		:comment_count
	)RETURNING id;
`
const updateBlog = `
	UPDATE blogs 
	SET
		updated_at =:updated_at,
		title =:title,
		description =:description
	WHERE 
		id =:id
	RETURNING *;
`
func (s *Storage) WriteBlog(ctx context.Context, b storage.Blog) (int64, error) {
	stmt, err := s.db.PrepareNamed(writeBlog)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, b); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) DeleteBlog(ctx context.Context, id int64, author_id int64)  error {
	var b storage.Blog
	if err:=s.db.Get(&b,"DELETE FROM blogs WHERE id=$1 RETURNING *",id);err!=nil{
		return err
	}
	return nil;
}

func (s *Storage) ReadBlog(ctx context.Context, id int64, author_id int64)  (*storage.Blog,bool, error) {
	var b storage.Blog
	if err:=s.db.Get(&b,"SELECT * FROM blogs WHERE id=$1",id);err!=nil{
		return nil,false,err;
	}
	if b.AuthorID != author_id{
		return &b,false,nil
	}
	return &b,true,nil;
}

func (s *Storage) ReadAllBlog(ctx context.Context)  ([]*storage.Blog,error) {
	blogs := []*storage.Blog{}
	if err :=s.db.Select(&blogs, "SELECT * FROM blogs");err!=nil{
		return []*storage.Blog{},err
	}
	return blogs,nil;
}

func (s *Storage) EditBlog(ctx context.Context, b storage.Blog)  (*storage.Blog, error) {
	stmt,err := s.db.PrepareNamed(updateBlog)
	if err != nil {
		return nil, err
	}
	if err := stmt.Get(&b, b); err != nil {
		return nil, err
	}
	return &b, nil
}
	