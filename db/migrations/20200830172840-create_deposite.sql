
-- +migrate Up
CREATE TABLE IF NOT EXISTS deposite (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT NOT NULL
);

-- +migrate Down
DROP TABLE deposite;