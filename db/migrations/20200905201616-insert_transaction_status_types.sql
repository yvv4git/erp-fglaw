
-- +migrate Up
INSERT INTO transaction_status_types
    (name)
VALUES
    ('DEPOSIT'),
    ('DELIVERED'),
    ('SOLD');

-- +migrate Down
DELETE FROM transaction_status_types;