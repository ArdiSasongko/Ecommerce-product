// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: product.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

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
