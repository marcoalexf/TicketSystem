-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.services (
    email CHARACTER VARYING NOT NULL,
    ipAddress CHARACTER VARYING NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (email, ipAddress)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE services;
-- +goose StatementEnd
