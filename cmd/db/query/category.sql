-- name: CreateCategory :one
INSERT INTO categories (name) VALUES ($1) RETURNING name;

-- name: GetCategory :one
SELECT id, name FROM categories WHERE name = $1;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE name = $1;

-- name: UpdateCategory :one
UPDATE categories SET name = $1 WHERE name = $2 RETURNING name;