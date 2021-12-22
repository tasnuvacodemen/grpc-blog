package postgres

import (
	"context"
	"reflect"
	"testing"

	"github.com/tasnuvatina/grpc-blog/todo/storage"
)

func TestStorage_Register(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.User
		want    int64
		wantErr bool
	}{
		{
			name: "REGISTER USER SUCCESS",
			in: storage.User{
				UserName: "username1",
				Email:    "username1@gmail.com",
				Password: "123456!",
			},
			want: 1,
		},
		{
			name: "REGISTER USER DUPLICATE TEST",
			in: storage.User{
				UserName: "username1",
				Email:    "username1@gmail.com",
				Password: "123456!",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Register(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetUser(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      string
		want    *storage.User
		wantErr bool
	}{
		{
			name: "GET TODO SUCCESS",
			in: "username1",
			want: &storage.User{
				ID: 1,
				UserName: "username1",
				Email:    "username1@gmail.com",
				Password: "123456!",
			},
		},
		{
			name: "GET TODO FAilure",
			in: "username1000",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
			got, err := s.GetUser(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
