-- name: InsertProductVariant :one
INSERT INTO product_variants (product_id, color, size, quantity) VALUES ($1, $2, $3, $4) RETURNING color, size, quantity;

-- name: UpdateProductVariant :one
UPDATE product_variants SET color = $1, size = $2, quantity = $3, updated_at = NOW() WHERE id = $4 AND product_id = $5 RETURNING color, size, quantity, updated_at;

-- name: GetVariantByID :one
SELECT id, product_id, color, size, quantity, created_at, updated_at FROM product_variants WHERE id = $1 AND product_id = $2;

-- name: GetVariantsByProductID :many
SELECT id, product_id, color, size, quantity, created_at, updated_at FROM product_variants WHERE product_id = $1 ORDER BY created_at DESC;

-- name: DeleteProductVariant :exec
DELETE FROM product_variants WHERE id = $1 AND product_id = $2;

-- name: CountVariantsByProductID :one
SELECT COUNT(id) FROM product_variants WHERE product_id = $1;