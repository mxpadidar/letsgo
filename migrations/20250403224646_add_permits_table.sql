-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS auth.permits (
    id UUID PRIMARY KEY,
    user_id INTEGER NOT NULL,
    role INTEGER NOT NULL,
    issued_at TIMESTAMP WITH TIME ZONE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES auth.users(id) ON DELETE CASCADE
);

-- Add indexes
CREATE INDEX IF NOT EXISTS idx_permits_user_id ON auth.permits(user_id);
CREATE INDEX IF NOT EXISTS idx_permits_issued_at ON auth.permits(issued_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS auth.permits;
-- +goose StatementEnd
