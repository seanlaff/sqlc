// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package querytest

import (
	"context"
)

const listBar = `-- name: ListBar :many
SELECT id_old, ip_old FROM bar_old
`

func (q *Queries) ListBar(ctx context.Context) ([]BarNew, error) {
	rows, err := q.db.Query(ctx, listBar)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BarNew
	for rows.Next() {
		var i BarNew
		if err := rows.Scan(&i.IDNew, &i.IpOld); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listFoo = `-- name: ListFoo :many
SELECT id_old as foo_old, id_old as baz_old
FROM bar_old
WHERE ip_old = $1 AND id_old = $2
`

type ListFooParams struct {
	IpOld IPProtocol
	IDNew int32
}

type ListFooRow struct {
	FooNew int32
	BazOld int32
}

func (q *Queries) ListFoo(ctx context.Context, arg ListFooParams) ([]ListFooRow, error) {
	rows, err := q.db.Query(ctx, listFoo, arg.IpOld, arg.IDNew)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListFooRow
	for rows.Next() {
		var i ListFooRow
		if err := rows.Scan(&i.FooNew, &i.BazOld); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
