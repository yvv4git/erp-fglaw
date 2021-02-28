
-- +migrate Up
CREATE TABLE IF NOT EXISTS suppliers (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    cuit_suppliers VARCHAR(255),
    Name           VARCHAR(50),
	Address        VARCHAR(255),
	Phone          VARCHAR(30)
);

-- +migrate Down
DROP TABLE suppliers;