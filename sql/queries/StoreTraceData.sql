-- name: StoreTraceData :exec
INSERT INTO traces (sn, hu, ref_list, placed)
VALUES (
    ?, ?, ?, ?
)
RETURNING *;