
-- +migrate Up
CREATE TABLE IF NOT EXISTS product_types (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    type_name VARCHAR(30) NOT NULL
);

-- +migrate Down
DROP TABLE product_types;