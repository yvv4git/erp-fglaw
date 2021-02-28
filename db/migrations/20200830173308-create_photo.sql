
-- +migrate Up
CREATE TABLE IF NOT EXISTS photo (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    file_name VARCHAR(255) NOT NULL,
    file_type_id INTEGER
);

-- +migrate Down
DROP TABLE photo;