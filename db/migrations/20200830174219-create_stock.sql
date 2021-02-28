
-- +migrate Up
CREATE TABLE IF NOT EXISTS stock (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    series_number    VARCHAR(255),
	product_id       INTEGER,
	expiration_date  INTEGER,
	entry_date       INTEGER,
	invoice_purchase INTEGER,
	invoice_sell     INTEGER,
	transaction_id   INTEGER
);

-- +migrate Down
DROP TABLE stock;