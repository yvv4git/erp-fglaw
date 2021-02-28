
-- +migrate Up
CREATE TABLE IF NOT EXISTS clients (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    number VARCHAR(100) NOT NULL,
    address VARCHAR(255),
    cuit_customer VARCHAR(255),
    client_phone VARCHAR(50),
    client_type_id INTEGER
);

-- +migrate Down
DROP TABLE clients;