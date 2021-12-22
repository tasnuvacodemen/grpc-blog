package postgres

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tasnuvatina/grpc-blog/todo/storage"
)

func TestCreateTodo(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      storage.Todo
		want    int64
		wantErr bool
	}{
		{
			name: "CREATE_TODO_SUCCESS",
			in: storage.Todo{
				Title: "this is title",
				Description: "this is description",
			},
			want: 1,
		},
		{
			name: "FAILED_DUPLICATE_TITLE",
			in: storage.Todo{
				Title: "this is title",
				Description: "this is description",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Create(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Create() = %v, want %v", got, tt.want)
			}

			
		})
	}
}

func TestGetTodo(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      int64
		want    *storage.Todo
		wantErr bool
	}{
		{
			name: "GET_TODO_SUCCESS",
			in: 1,
			want: &storage.Todo{
				ID:          1,
				Title:       "this is title",
				Description: "this is description",
				IsCompleted: false,
			},
		},
		{
			name: "FAILED_TODO_DOES_NOT_EXIST",
			in: 100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Get(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("Diff: got -, want += %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestUpdateTodo(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      storage.Todo
		want    *storage.Todo
		wantErr bool
	}{
		{
			name: "UPDATE_TODO_SUCCESS",
			in: storage.Todo{
				ID: 1,
				Title: "this is title update",
				Description: "this is description update",
			},
			want: &storage.Todo{
				ID: 1,
				Title: "this is title update",
				Description: "this is description update",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Update(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Diff: got -, want += %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestDeleteTodo(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      int64
		wantErr bool
	}{
		{
			name: "DELETE_TODO_SUCCESS",
			in: 1,
		},
		{
			name: "FAILED_TO_DELETE_TODO_ID_DOES_NOT_EXIST",
			in: 100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := s.Delete(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}