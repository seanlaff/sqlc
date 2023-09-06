// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package querytest

import (
	"context"
)

const schemaScopedDelete = `-- name: SchemaScopedDelete :exec
DELETE FROM foo.bar WHERE id = ?
`

func (q *Queries) SchemaScopedDelete(ctx context.Context, id uint64) error {
	_, err := q.db.ExecContext(ctx, schemaScopedDelete, id)
	return err
}
