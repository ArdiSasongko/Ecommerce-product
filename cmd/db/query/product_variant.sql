-- name: InsertProductVariant :one
INSERT INTO product_variants (product_id, color, size, quantity) VALUES ($1, $2, $3, $4) RETURNING color, size, quantity;