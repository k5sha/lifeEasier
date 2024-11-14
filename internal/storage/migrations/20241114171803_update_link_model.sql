-- +goose Up
-- +goose StatementBegin
ALTER TABLE links
    ADD COLUMN scheduled_at TIMESTAMP NOT NULL ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE links
    DROP COLUMN scheduled_at ;
-- +goose StatementEnd
