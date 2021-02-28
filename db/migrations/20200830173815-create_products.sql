
-- +migrate Up
CREATE TABLE IF NOT EXISTS products (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    suppliers_id INTEGER,
    product_type_id INTEGER NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    product_brand VARCHAR(255) NOT NULL,
    product_photo_id INTEGER
);

-- +migrate Down
DROP TABLE products;