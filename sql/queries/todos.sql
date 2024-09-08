-- name: Save :one
INSERT INTO todos (title, status)
VALUES ($1, $2)
RETURNING *;

-- name: FindById :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: FindAll :many
SELECT * FROM todos;

-- name: DeleteById :exec
DELETE FROM todos WHERE
id = $1;

-- name: UpdateById :one
UPDATE todos
SET title = $1, status = $2, updated_at = NOW()
WHERE id = $3
RETURNING *;
