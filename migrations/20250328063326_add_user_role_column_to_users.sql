-- +goose Up
-- +goose StatementBegin
ALTER TABLE auth.users
ADD COLUMN role INTEGER NOT NULL DEFAULT 2;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE auth.users
DROP COLUMN role;
-- +goose StatementEnd
