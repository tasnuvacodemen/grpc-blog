-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS comments (
    id SERIAL NOT NULL,
    blog_id INT NOT NULL,
    user_id INT NOT NULL,
    user_name TEXT NOT NULL,
    content TEXT NOT NULL,
    commented_at TEXT NOT NULL,
    
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comments ;
-- +goose StatementEnd
