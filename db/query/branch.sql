-- name: CreateBranch :one
INSERT INTO branches(
    name,
    "maxUsers"
) VALUES (
             $1, $2
         ) RETURNING *;

-- name: GetBranchById :one
SELECT * FROM branches
WHERE id = $1 LIMIT 1;

-- name: ListBranches :many
SELECT * FROM branches
ORDER BY id
LIMIT $1
    OFFSET $2;

-- name: UpdateBranch :one
UPDATE branches
SET
    name = coalesce(sqlc.narg('name'), name),
    "maxUsers" = coalesce(sqlc.narg('maxUsers'), "maxUsers")
WHERE id = sqlc.arg('branch_id')
RETURNING *;

-- name: DeleteBranch :exec
DELETE FROM branches
WHERE id = $1;

-- name: DeleteBranches :exec
DELETE FROM branches
WHERE id = ANY($1::int[]);