-- +goose Up
-- +goose StatementBegin
CREATE TABLE links (
        id SERIAL PRIMARY KEY,
        link VARCHAR(255) NOT NULL,
        message VARCHAR(255),
        used BOOLEAN DEFAULT FALSE,
        user_id INTEGER,
        created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Add the foreign key reference to Users table
ALTER TABLE links
    ADD CONSTRAINT fk_links_user_id
        FOREIGN KEY (user_id)
            REFERENCES Users (id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS links;
-- +goose StatementEnd
