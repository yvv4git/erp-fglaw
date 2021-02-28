
-- +migrate Up
CREATE TABLE IF NOT EXISTS transaction_status_types (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(20)
);

-- +migrate Down
DROP TABLE transaction_status_types;