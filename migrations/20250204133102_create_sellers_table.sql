-- +goose Up
-- +goose StatementBegin
CREATE TABLE sellers (
                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         name VARCHAR(100) NOT NULL,
                         phone VARCHAR(20) NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sellers;
-- +goose StatementEnd
