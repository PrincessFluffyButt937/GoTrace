-- name: StoreBoardData :one
INSERT INTO boards (sn, pb, rev)
VALUES (
    ?, ?, ?
)
RETURNING *;