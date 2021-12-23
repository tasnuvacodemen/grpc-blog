package postgres

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tasnuvatina/grpc-blog/todo/storage"
)

func TestStorage_WriteBlog(t *testing.T) {
	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      storage.Blog
		want    int64
		wantErr bool
	}{
		{
			name: "CREATE_BLOG_SUCCESS",
			in: storage.Blog{
				AuthorID:      23,
				AuthorName:    "tasnuva",
				CreatedAt:     "test created",
				UpdateAt:      "test updated",
				PictureString: "test picture",
				Title:         "test title",
				Description:   "test description",
				UpvoteCount:   0,
				DownvoteCount: 0,
				CommentsCount: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.WriteBlog(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.WriteBlog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.WriteBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_ReadBlog(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name      string
		id        int64
		author_id int64
		want      *storage.Blog
		want1     bool
		wantErr   bool
	}{
		{
			name:      "READ BLOG SUCCESS",
			id:        1,
			author_id: 23,
			want: &storage.Blog{
				ID:            1,
				AuthorID:      23,
				AuthorName:    "tasnuva",
				CreatedAt:     "test created",
				UpdateAt:      "test updated",
				PictureString: "test picture",
				Title:         "test title",
				Description:   "test description",
				UpvoteCount:   0,
				DownvoteCount: 0,
				CommentsCount: 0,
			},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, got1, err := s.ReadBlog(context.TODO(), tt.id, tt.author_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.ReadBlog() error = %v, wantErr %v", err, tt.wantErr)
				t.Errorf("Diff: got -, want += %v", cmp.Diff(err, tt.wantErr))
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.ReadBlog() got += %#v, want - %#v", got, tt.want)
				t.Errorf("Diff: got -, want += %v", cmp.Diff(got, tt.want))
			}
			if got1 != tt.want1 {
				t.Errorf("Storage.ReadBlog() got1 += %#v, want - %#v", got1, tt.want1)
				t.Errorf("Diff: got -, want += %v", cmp.Diff(got1, tt.want1))
			}
		})
	}
}

func TestStorage_ReadAllBlog(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		want    []*storage.Blog
		wantErr bool
	}{
		{
			name: "READ ALL BLOG SUCCESS",
			want: []*storage.Blog{
				{
					ID:            1,
					AuthorID:      23,
					AuthorName:    "tasnuva",
					CreatedAt:     "test created",
					UpdateAt:      "test updated",
					PictureString: "test picture",
					Title:         "test title",
					Description:   "test description",
					UpvoteCount:   0,
					DownvoteCount: 0,
					CommentsCount: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ReadAllBlog(context.TODO())
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.ReadAllBlog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.ReadAllBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_EditBlog(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Blog
		want    *storage.Blog
		wantErr bool
	}{
		{
			name: "EDIT BLOG SUCCESS",
			in: storage.Blog{
				ID:            1,
				AuthorID:      23,
				AuthorName:    "tasnuva",
				CreatedAt:     "test edited",
				UpdateAt:      "test edited",
				PictureString: "test picture",
				Title:         "test title edited",
				Description:   "test description edited",
				UpvoteCount:   1,
				DownvoteCount: 2,
				CommentsCount: 3,
			},
			want: &storage.Blog{
				ID:            1,
				AuthorID:      23,
				AuthorName:    "tasnuva",
				CreatedAt:     "test created",
				UpdateAt:      "test edited",
				PictureString: "test picture",
				Title:         "test title edited",
				Description:   "test description edited",
				UpvoteCount:   0,
				DownvoteCount: 0,
				CommentsCount: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.EditBlog(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.EditBlog() error = %v, wantErr %v", err, tt.wantErr)
				t.Errorf("Diff: got -, want += %v", cmp.Diff(err, tt.wantErr))
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.EditBlog() = %v, want %v", got, tt.want)
				t.Errorf("Diff: got -, want += %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestStorage_DeleteBlog(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name     string
		id       int64
		autherId int64
		wantErr  bool
	}{
		{
			name:     "DELETE BLOG SUCCESS",
			id:       1,
			autherId: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.DeleteBlog(context.TODO(), tt.id, tt.autherId); (err != nil) != tt.wantErr {
				t.Errorf("Storage.DeleteBlog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_UpvoteBlog(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Upvote
		want    int64
		wantErr bool
	}{
		{
			name: "UPVOTE SUCCESS",
			in: storage.Upvote{
				BlogID: 2,
				UserID: 24,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.UpvoteBlog(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.UpvoteBlog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.UpvoteBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_DownVoteBlog(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Downvote
		want    int64
		wantErr bool
	}{
		{
			name: "DOWNVOTE SUCCESS",
			in: storage.Downvote{
				BlogID: 2,
				UserID: 24,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.DownVoteBlog(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.DownVoteBlog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.DownVoteBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_CommentBlog(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Comment
		want    int64
		wantErr bool
	}{
		{
			name: "DOWNVOTE SUCCESS",
			in: storage.Comment{
				BlogID:      2,
				UserID:      24,
				UserName:    "Test user",
				Content:     "test content",
				CommentedAt: "test date",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CommentBlog(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.CommentBlog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.CommentBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetUpvote(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		blog_id int64
		user_id int64
		want    *storage.Upvote
		want1   int64
		wantErr bool
	}{
		{
			name:    "GET ONE UPVOTE SUCCESS",
			blog_id: 2,
			user_id: 24,
			want: &storage.Upvote{
				ID:     1,
				BlogID: 2,
				UserID: 24,
			},
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := s.GetUpvote(context.TODO(), tt.blog_id, tt.user_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetUpvote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetUpvote() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Storage.GetUpvote() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStorage_GetDownvote(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		blog_id int64
		user_id int64
		want    *storage.Downvote
		want1   int64
		wantErr bool
	}{
		{
			name:    "GET ONE DOWN SUCCESS",
			blog_id: 2,
			user_id: 24,
			want: &storage.Downvote{
				ID:     1,
				BlogID: 2,
				UserID: 24,
			},
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := s.GetDownvote(context.TODO(), tt.blog_id, tt.user_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetDownvote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetDownvote() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Storage.GetDownvote() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStorage_GetAllComments(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		blog_id int64
		want    []*storage.Comment
		wantErr bool
	}{
		{
			name:    "Get all comments success",
			blog_id: 2,
			want: []*storage.Comment{
				{
					ID:          1,
					BlogID:      2,
					UserID:      24,
					UserName:    "Test user",
					Content:     "test content",
					CommentedAt: "test date",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GetAllComments(context.TODO(), tt.blog_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetAllComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetAllComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetAllUpvote(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		blog_id int64
		want    []*storage.Upvote
		wantErr bool
	}{
		{
			name:    "GET ALL UPVOTES SUCCESS",
			blog_id: 2,
			want: []*storage.Upvote{
				{
					ID:     1,
					BlogID: 2,
					UserID: 24,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GetAllUpvote(context.TODO(), tt.blog_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetAllUpvote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetAllUpvote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetAllDownvote(t *testing.T) {
	s :=newTestStorage(t)
	tests := []struct {
		name    string
		blog_id int64
		want    []*storage.Downvote
		wantErr bool
	}{
		{
			name:    "GET ALL DOWNVOTES SUCCESS",
			blog_id: 2,
			want: []*storage.Downvote{
				{
					ID:     1,
					BlogID: 2,
					UserID: 24,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
			got, err := s.GetAllDownvote(context.TODO(), tt.blog_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetAllDownvote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetAllDownvote() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestStorage_RevertUpvoteBlog(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name      string
		upvote_id int64
		user_id   int64
		wantErr   bool
	}{
		{
			name:      "DELETE UPVOTE SUCCESS",
			upvote_id: 1,
			user_id:   24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := s.RevertUpvoteBlog(context.TODO(), tt.upvote_id, tt.user_id); (err != nil) != tt.wantErr {
				t.Errorf("Storage.RevertUpvoteBlog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_RevertDownVoteBlog(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name        string
		downvote_id int64
		user_id     int64
		wantErr     bool
	}{
		{
			name:        "DELETE DOWNVOTE SUCCESS",
			downvote_id: 1,
			user_id:     24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := s.RevertDownVoteBlog(context.TODO(), tt.downvote_id, tt.user_id); (err != nil) != tt.wantErr {
				t.Errorf("Storage.RevertDownVoteBlog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

