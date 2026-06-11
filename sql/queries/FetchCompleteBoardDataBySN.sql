-- name: FetchBoardBySN :many
SELECT * FROM traces
INNER JOIN boards ON traces.sn = boards.sn
INNER JOIN comps ON traces.hu = comps.hu
WHERE traces.sn = ?;