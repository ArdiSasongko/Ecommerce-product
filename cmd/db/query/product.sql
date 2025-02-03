-- name: InsertProduct :one
INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id;

-- name: UpdateProduct :one
UPDATE products SET name = $1, description = $2, price = $3, updated_at = NOW() WHERE id = $4 RETURNING id, name, description, price, updated_at;

-- name: GetProduct :one
SELECT id, name, description, price, created_at, updated_at FROM products WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;

-- name: ListProducts :many
SELECT id, name, description, price, created_at, updated_at FROM products ORDER BY created_at DESC;

-- name: CountProducts :one
SELECT COUNT(id) FROM products;

-- name: GetProductDetails :one
SELECT
    p.*,
    ARRAY_AGG(c.name) AS categories
FROM products p
LEFT JOIN product_categories pc ON p.id = pc.product_id
LEFT JOIN categories c ON pc.category_id = c.id
WHERE p.id = $1
GROUP BY p.id;
