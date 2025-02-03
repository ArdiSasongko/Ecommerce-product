// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: product.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countProducts = `-- name: CountProducts :one
SELECT COUNT(id) FROM products
`

func (q *Queries) CountProducts(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countProducts)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT id, name, description, price, created_at, updated_at FROM products WHERE id = $1
`

func (q *Queries) GetProduct(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProductDetails = `-- name: GetProductDetails :one
SELECT
    p.id, p.name, p.description, p.price, p.created_at, p.updated_at,
    ARRAY_AGG(c.name) AS categories
FROM products p
LEFT JOIN product_categories pc ON p.id = pc.product_id
LEFT JOIN categories c ON pc.category_id = c.id
WHERE p.id = $1
GROUP BY p.id
`

type GetProductDetailsRow struct {
	ID          int32
	Name        string
	Description pgtype.Text
	Price       pgtype.Numeric
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
	Categories  interface{}
}

func (q *Queries) GetProductDetails(ctx context.Context, id int32) (GetProductDetailsRow, error) {
	row := q.db.QueryRow(ctx, getProductDetails, id)
	var i GetProductDetailsRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Categories,
	)
	return i, err
}

const insertProduct = `-- name: InsertProduct :one
INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id
`

type InsertProductParams struct {
	Name        string
	Description pgtype.Text
	Price       pgtype.Numeric
}

func (q *Queries) InsertProduct(ctx context.Context, arg InsertProductParams) (int32, error) {
	row := q.db.QueryRow(ctx, insertProduct, arg.Name, arg.Description, arg.Price)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const listProducts = `-- name: ListProducts :many
SELECT id, name, description, price, created_at, updated_at FROM products ORDER BY created_at DESC
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products SET name = $1, description = $2, price = $3, updated_at = NOW() WHERE id = $4 RETURNING id, name, description, price, updated_at
`

type UpdateProductParams struct {
	Name        string
	Description pgtype.Text
	Price       pgtype.Numeric
	ID          int32
}

type UpdateProductRow struct {
	ID          int32
	Name        string
	Description pgtype.Text
	Price       pgtype.Numeric
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (UpdateProductRow, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.ID,
	)
	var i UpdateProductRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.UpdatedAt,
	)
	return i, err
}
