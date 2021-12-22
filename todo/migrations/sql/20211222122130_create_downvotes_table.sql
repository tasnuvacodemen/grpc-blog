-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS downvotes (
    id SERIAL NOT NULL,
    blog_id INT NOT NULL,
    user_id INT NOT NULL,
    
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS downvotes ;
-- +goose StatementEnd
