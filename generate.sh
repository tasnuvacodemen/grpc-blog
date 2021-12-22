protoc --go_out=plugins=grpc:. proto/todo/todo.proto 

go run migrations/migrate.go create create_todos_table sql 

go run migrations/migrate.go up

grpcurl -plaintext 127.0.0.1:3000 list todo.TodoService 


DATABASE_CONNECTION="user=postgres password=password host=localhost port=5432 sslmode=disable" go test