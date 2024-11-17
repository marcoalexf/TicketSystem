-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.services (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    email CHARACTER VARYING NOT NULL,
    ipAddress CHARACTER VARYING NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id, email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE services;
-- +goose StatementEnd
