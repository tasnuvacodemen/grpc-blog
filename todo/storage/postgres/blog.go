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
// upvote query
const upvote =`
	INSERT INTO upvotes(
		blog_id,
		user_id
	) VALUES(
		:blog_id,
		:user_id
	)RETURNING id;
`

// downvote query
const downvote =`
	INSERT INTO downvotes(
		blog_id,
		user_id
	) VALUES(
		:blog_id,
		:user_id
	)RETURNING id;
`
// downvote query
const comment =`
	INSERT INTO comments(
		blog_id,
		user_id,
		user_name,
		content,
		commented_at
	) VALUES(
		:blog_id,
		:user_id,
		:user_name,
		:content,
		:commented_at
	)RETURNING id;
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

func (s *Storage) ReadAllSearchedBlog(ctx context.Context,search_value string)  ([]*storage.Blog,error) {
	blogs := []*storage.Blog{}
	if err :=s.db.Select(&blogs, "SELECT * FROM blogs WHERE (blogs.author_name ILIKE '%' || $1 || '%' OR blogs.title ILIKE '%' || $1 || '%' OR blogs.description ILIKE '%' || $1 || '%')",search_value);err!=nil{
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


// upvote functions
func (s *Storage) UpvoteBlog(ctx context.Context, u storage.Upvote) (int64, error) {
	stmt, err := s.db.PrepareNamed(upvote)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, u); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) GetUpvote(ctx context.Context, blog_id int64, user_id int64)  (*storage.Upvote,int64, error) {
	var b storage.Upvote
	if err:=s.db.Get(&b,"SELECT * FROM upvotes WHERE blog_id=$1 AND user_id=$2",blog_id,user_id);err!=nil{
		return nil,0,err;
	}
	return &b,b.ID,nil;
}

func (s *Storage) GetAllUpvote(ctx context.Context, blog_id int64)  ([]*storage.Upvote, error) {
	upvotes := []*storage.Upvote{}
	if err :=s.db.Select(&upvotes, "SELECT * FROM upvotes WHERE blog_id=$1",blog_id);err!=nil{
		return []*storage.Upvote{},err
	}
	return upvotes,nil;
}

func (s *Storage)GetAllUpvoteCount(ctx context.Context, blog_id int64)(int64,error)  {
	var upvoteCount int64
	if err :=s.db.Get(&upvoteCount,"SELECT COUNT(id) FROM upvotes WHERE blog_id=$1",blog_id);err!=nil{
		return 0,err

	}
	return upvoteCount,nil
}

func (s *Storage) RevertUpvoteBlog(ctx context.Context, upvote_id int64, user_id int64)  error {
	var b storage.Upvote
	if err:=s.db.Get(&b,"DELETE FROM upvotes WHERE id=$1 RETURNING *",upvote_id);err!=nil{
		return err
	}
	return nil;
}

// downvote functions
func (s *Storage) DownVoteBlog(ctx context.Context, u storage.Downvote) (int64, error) {
	stmt, err := s.db.PrepareNamed(downvote)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, u); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) GetDownvote(ctx context.Context, blog_id int64, user_id int64)  (*storage.Downvote,int64, error) {
	var b storage.Downvote
	if err:=s.db.Get(&b,"SELECT * FROM downvotes WHERE blog_id=$1 AND user_id=$2",blog_id,user_id);err!=nil{
		return nil,0,err;
	}
	return &b,b.ID,nil;
}
func (s *Storage) GetAllDownvote(ctx context.Context, blog_id int64)  ([]*storage.Downvote, error) {
	downvotes := []*storage.Downvote{}
	if err :=s.db.Select(&downvotes, "SELECT * FROM downvotes WHERE blog_id=$1",blog_id);err!=nil{
		return []*storage.Downvote{},err
	}
	return downvotes,nil;
}
func (s *Storage)GetAllDownvoteCount(ctx context.Context, blog_id int64)(int64,error)  {
	var downvoteCount int64
	if err :=s.db.Get(&downvoteCount,"SELECT COUNT(id) FROM downvotes WHERE blog_id=$1",blog_id);err!=nil{
		return 0,err

	}
	return downvoteCount,nil
}

func (s *Storage) RevertDownVoteBlog(ctx context.Context, downvote_id int64, user_id int64)  error {
	var b storage.Downvote
	if err:=s.db.Get(&b,"DELETE FROM downvotes WHERE id=$1 RETURNING *",downvote_id);err!=nil{
		return err
	}
	return nil;
}

// Comment functions
func (s *Storage) CommentBlog(ctx context.Context, u storage.Comment) (int64, error) {
	stmt, err := s.db.PrepareNamed(comment)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, u); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) GetAllComments(ctx context.Context, blog_id int64)  ([]*storage.Comment, error) {
	comments := []*storage.Comment{}
	if err :=s.db.Select(&comments, "SELECT * FROM comments WHERE blog_id=$1",blog_id);err!=nil{
		return []*storage.Comment{},err
	}
	return comments,nil;
}
func (s *Storage)GetAllCommentCount(ctx context.Context, blog_id int64)(int64,error)  {
	var commentsCount int64
	if err :=s.db.Get(&commentsCount,"SELECT COUNT(id) FROM comments WHERE blog_id=$1",blog_id);err!=nil{
		return 0,err

	}
	return commentsCount,nil
}
	