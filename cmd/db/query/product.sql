-- name: InsertProduct :one
INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id;

-- name: UpdateProduct :one
UPDATE products SET name = $1, description = $2, price = $3, updated_at = NOW() WHERE id = $4 RETURNING id, name, description, price, updated_at;

-- name: GetProduct :one
SELECT id, name, description, price, created_at, updated_at FROM products WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;
