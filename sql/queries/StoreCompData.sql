-- name: StoreCompData :one
INSERT INTO comps (hu, pn, lot)
VALUES (
    ?, ?, ?
)
RETURNING *;