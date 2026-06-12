-- name: FetchSnContainingHandlingUnit :many
SELECT DISTINCT sn FROM traces
WHERE hu = ?;