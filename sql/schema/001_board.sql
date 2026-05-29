-- +goose Up
CREATE TABLE boards (
    sn TEXT NOT NULL PRIMARY KEY,
    pb TEXT NOT NULL,
    rev TEXT NOT NULL
);

-- +goose Down
DROP TABLE boards;