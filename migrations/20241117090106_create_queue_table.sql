-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.queues (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    position INT NOT NULL,
    service_id uuid NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (service_id) REFERENCES services (id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.queues;
-- +goose StatementEnd
