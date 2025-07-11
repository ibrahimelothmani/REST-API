-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    users (
        id BIGSERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password_hash VARCHAR(255) NOT NULL,
        bio TEXT,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    )
    -- +goose StatementEnd
    -- +goose Down
    -- +goose StatementBegin
DROP TABLE users;

-- +goose StatementEnd