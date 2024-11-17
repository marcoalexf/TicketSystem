-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.queues (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    position INT NOT NULL,
    service_id uuid NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (service_id) REFERENCES services (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT unique_service_position UNIQUE (service_id, position)
);

CREATE OR REPLACE FUNCTION increment_position()
RETURNS TRIGGER AS $$
DECLARE
    max_position INT;
BEGIN
    -- Get the max position for the given service_id, defaulting to -1 if none exist
    SELECT COALESCE(MAX(position), -1) + 1 INTO max_position
    FROM queues
    WHERE service_id = NEW.service_id;

    -- Assign the computed position to the NEW record
    NEW.position := max_position;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_position
BEFORE INSERT ON queues
FOR EACH ROW
EXECUTE FUNCTION increment_position();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.queues;
-- +goose StatementEnd
