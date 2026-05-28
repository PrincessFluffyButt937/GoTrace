-- +goose Up
CREATE TABLE trace (
    id INT NOT NULL PRIMARY KEY,
    sn TEXT NOT NULL REFERENCES board(sn),
    hu TEXT NOT NULL REFERENCES comp(hu),
    ref_list TEXT NOT NULL,
    placed TIMESTAMP NOT NULL,
);

-- +goose Down
DROP TABLE trace;