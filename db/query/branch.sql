-- name: CreateBranch :one
INSERT INTO "Branches"(
    "Name",
    "MaxUsers"
) VALUES (
             $1, $2
         ) RETURNING *;

-- name: GetBranchById :one
SELECT * FROM "Branches"
WHERE "Id" = $1 LIMIT 1;

-- name: ListBranches :many
SELECT * FROM "Branches"
ORDER BY "Id"
LIMIT $1
    OFFSET $2;

-- name: UpdateBranch :one
UPDATE "Branches"
SET
    "Name" = coalesce(sqlc.narg('name'), "Name"),
    "MaxUsers" = coalesce(sqlc.narg('maxUsers'), "MaxUsers")
WHERE "Id" = sqlc.arg('branch_id')
RETURNING *;

-- name: DeleteBranch :exec
DELETE FROM "Branches"
WHERE "Id" = $1;

-- name: DeleteBranches :exec
DELETE FROM "Branches"
WHERE "Id" = ANY($1::int[]);