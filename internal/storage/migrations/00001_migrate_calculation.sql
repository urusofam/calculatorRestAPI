-- +goose Up
-- +goose StatementBegin
CREATE TABLE calculations (
    id varchar(100) NOT NULL UNIQUE,
    expression varchar(50) NOT NULL,
    result varchar(50) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE calculations;
-- +goose StatementEnd
