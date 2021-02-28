
-- +migrate Up
CREATE TABLE IF NOT EXISTS client_types (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    client_type VARCHAR(255) NOT NULL,
    acting_as VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE client_types;