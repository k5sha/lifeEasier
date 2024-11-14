-- +goose Up
-- +goose StatementBegin
CREATE TABLE links (
        id SERIAL PRIMARY KEY,
        link VARCHAR(1024) NOT NULL,
        message VARCHAR(255),
        chat_id INTEGER,
        posted_at TIMESTAMP,
        created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS links;
-- +goose StatementEnd
