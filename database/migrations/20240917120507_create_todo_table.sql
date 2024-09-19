-- +goose Up
-- +goose StatementBegin
CREATE TYPE task_status AS ENUM ('PENDING', 'COMPLETED');

CREATE TABLE todo (
    id SERIAL PRIMARY KEY,
    task TEXT,
    status task_status,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todo;
DROP TYPE task_status;
-- +goose StatementEnd
