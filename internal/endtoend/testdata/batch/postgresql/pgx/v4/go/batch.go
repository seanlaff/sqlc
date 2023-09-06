// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: batch.go

package querytest

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v4"
)

var (
	ErrBatchAlreadyClosed = errors.New("batch already closed")
)

const getValues = `-- name: GetValues :batchmany
SELECT a, b
FROM myschema.foo
WHERE b = $1
`

type GetValuesBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

func (q *Queries) GetValues(ctx context.Context, b []sql.NullInt32) *GetValuesBatchResults {
	batch := &pgx.Batch{}
	for _, a := range b {
		vals := []interface{}{
			a,
		}
		batch.Queue(getValues, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &GetValuesBatchResults{br, len(b), false}
}

func (b *GetValuesBatchResults) Query(f func(int, []MyschemaFoo, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		var items []MyschemaFoo
		if b.closed {
			if f != nil {
				f(t, items, ErrBatchAlreadyClosed)
			}
			continue
		}
		err := func() error {
			rows, err := b.br.Query()
			if err != nil {
				return err
			}
			defer rows.Close()
			for rows.Next() {
				var i MyschemaFoo
				if err := rows.Scan(&i.A, &i.B); err != nil {
					return err
				}
				items = append(items, i)
			}
			return rows.Err()
		}()
		if f != nil {
			f(t, items, err)
		}
	}
}

func (b *GetValuesBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}

const insertValues = `-- name: InsertValues :batchone
INSERT INTO myschema.foo (a, b)
VALUES ($1, $2)
RETURNING a
`

type InsertValuesBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

type InsertValuesParams struct {
	A sql.NullString
	B sql.NullInt32
}

func (q *Queries) InsertValues(ctx context.Context, arg []InsertValuesParams) *InsertValuesBatchResults {
	batch := &pgx.Batch{}
	for _, a := range arg {
		vals := []interface{}{
			a.A,
			a.B,
		}
		batch.Queue(insertValues, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &InsertValuesBatchResults{br, len(arg), false}
}

func (b *InsertValuesBatchResults) QueryRow(f func(int, sql.NullString, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		var a sql.NullString
		if b.closed {
			if f != nil {
				f(t, a, ErrBatchAlreadyClosed)
			}
			continue
		}
		row := b.br.QueryRow()
		err := row.Scan(&a)
		if f != nil {
			f(t, a, err)
		}
	}
}

func (b *InsertValuesBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}

const updateValues = `-- name: UpdateValues :batchexec
UPDATE myschema.foo SET a = $1, b = $2
`

type UpdateValuesBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

type UpdateValuesParams struct {
	A sql.NullString
	B sql.NullInt32
}

func (q *Queries) UpdateValues(ctx context.Context, arg []UpdateValuesParams) *UpdateValuesBatchResults {
	batch := &pgx.Batch{}
	for _, a := range arg {
		vals := []interface{}{
			a.A,
			a.B,
		}
		batch.Queue(updateValues, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &UpdateValuesBatchResults{br, len(arg), false}
}

func (b *UpdateValuesBatchResults) Exec(f func(int, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		if b.closed {
			if f != nil {
				f(t, ErrBatchAlreadyClosed)
			}
			continue
		}
		_, err := b.br.Exec()
		if f != nil {
			f(t, err)
		}
	}
}

func (b *UpdateValuesBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}
