package storage

type Todo struct{
	ID int64 `db:"id"`
	Title string `db:"title"`
	Description string `db:"description"`
	IsCompleted bool `db:"is_completed"`
}

type User struct{
	ID int64 `db:"id"`
	UserName string `db:"user_name"`
	Email string `db:"email"`
	Password string `db:"password"`
}

type Blog struct{
     ID int64  `db:"id"`
     AuthorID int64  `db:"author_id"`
     AuthorName string  `db:"author_name"`
     CreatedAt string  `db:"created_at"`
     UpdateAt string  `db:"updated_at"`
     PictureString string  `db:"picture_string"`
     Title  string `db:"title"`
     Description  string `db:"description"`
     UpvoteCount  int64 `db:"upvote_count"`
     DownvoteCount int64 `db:"downvote_count"`
     CommentsCount int64 `db:"comment_count"`
}
type Upvote struct{
     ID  int64 `db:"id"`
     BlogID int64 `db:"blog_id"`
     UserID int64 `db:"user_id"`
}
type Downvote struct{
	ID  int64 `db:"id"`
	BlogID int64 `db:"blog_id"`
	UserID int64 `db:"user_id"`
}
type Comment struct{
	ID  int64 `db:"id"`
     BlogID int64 `db:"blog_id"`
	UserID int64 `db:"user_id"`
     UserName  string `db:"user_name"`
     Content  string `db:"content"`
     CommentedAt string `db:"commented_at"`
}