-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS blogs (
    id SERIAL NOT NULL,
    author_id INT,
    author_name TEXT, 
    created_at TEXT NOT NULL,
    updated_at TEXT,
    picture_string TEXT,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    upvote_count INT,
    downvote_count INT,
    comment_count INT,

    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS blogs ;
-- +goose StatementEnd
