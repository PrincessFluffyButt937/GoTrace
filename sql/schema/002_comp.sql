-- +goose Up
CREATE TABLE comp (
    hu TEXT NOT NULL PRIMARY KEY,
    pn TEXT NOT NULL,
    lot TEXT NOT NULL,
);

-- +goose Down
DROP TABLE comp;