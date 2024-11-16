-- +goose StatementBegin
-- +goose Up
CREATE TABLE services (
    email CHARACTER VARYING NOT NULL,
    ipAddress CHARACTER VARYING NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (email, ipAddress)
);

-- +goose Down
DROP TABLE services;
-- +goose StatementEnd
