-- name: StoreTraceData :one
INSERT INTO traces (sn, hu, ref_list, placed)
VALUES (
    ?, ?, ?, ?
)
RETURNING *;