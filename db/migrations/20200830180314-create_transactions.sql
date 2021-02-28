
-- +migrate Up
CREATE TABLE IF NOT EXISTS transactions (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
	client_id INTEGER NOT NULL,
	deposite_id INTEGER,
	transaction_date INTEGER NOT NULL,
	transaction_status_type_id INTEGER NOT NULL,
	delivery_days INTEGER NOT NULL,
	product_id INTEGER NOT NULL,
	series_number VARCHAR(255)
);

-- +migrate Down
DROP TABLE transactions;