-- +goose Up
CREATE TABLE traces (
    id INTEGER PRIMARY KEY,
    sn TEXT NOT NULL REFERENCES boards(sn),
    hu TEXT NOT NULL REFERENCES comps(hu),
    ref_list TEXT NOT NULL,
    placed TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE traces;