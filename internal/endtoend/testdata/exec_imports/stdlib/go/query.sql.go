// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package querytest

import (
	"context"
)

const bar = `-- name: Bar :exec
SELECT bar
FROM foo
`

func (q *Queries) Bar(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, bar)
	return err
}

const bars = `-- name: Bars :exec
SELECT bars
FROM foo
`

func (q *Queries) Bars(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, bars)
	return err
}
