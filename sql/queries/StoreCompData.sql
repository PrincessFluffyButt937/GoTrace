-- name: StoreCompData :exec
INSERT INTO comps (hu, pn, lot)
VALUES (
    ?, ?, ?
)
RETURNING *;