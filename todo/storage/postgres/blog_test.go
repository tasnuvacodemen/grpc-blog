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
