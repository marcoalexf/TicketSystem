-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.services (
    id uuid NOT NULL UNIQUE DEFAULT gen_random_uuid(),
    email CHARACTER VARYING UNIQUE NOT NULL,
    ip_address CHARACTER VARYING NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id, email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE services;
-- +goose StatementEnd
