-- name: StoreBoardData :exec
INSERT INTO boards (sn, pb, rev)
VALUES (
    ?, ?, ?
)
RETURNING *;