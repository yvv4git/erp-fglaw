
-- +migrate Up
CREATE TABLE IF NOT EXISTS file_types (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    type_name VARCHAR(30) NOT NULL
);

-- +migrate Down
DROP TABLE file_types;