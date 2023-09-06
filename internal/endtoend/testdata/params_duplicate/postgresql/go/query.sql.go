// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const selectUserByID = `-- name: SelectUserByID :many
SELECT first_name from
users where ($1 = id OR $1 = 0)
`

func (q *Queries) SelectUserByID(ctx context.Context, id int32) ([]sql.NullString, error) {
	rows, err := q.db.QueryContext(ctx, selectUserByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var first_name sql.NullString
		if err := rows.Scan(&first_name); err != nil {
			return nil, err
		}
		items = append(items, first_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectUserByName = `-- name: SelectUserByName :many
SELECT first_name
FROM users
WHERE first_name = $1
   OR last_name = $1
`

func (q *Queries) SelectUserByName(ctx context.Context, name sql.NullString) ([]sql.NullString, error) {
	rows, err := q.db.QueryContext(ctx, selectUserByName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var first_name sql.NullString
		if err := rows.Scan(&first_name); err != nil {
			return nil, err
		}
		items = append(items, first_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectUserQuestion = `-- name: SelectUserQuestion :many
SELECT first_name from
users where ($1 = id OR  $1 = 0)
`

func (q *Queries) SelectUserQuestion(ctx context.Context, id int32) ([]sql.NullString, error) {
	rows, err := q.db.QueryContext(ctx, selectUserQuestion, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var first_name sql.NullString
		if err := rows.Scan(&first_name); err != nil {
			return nil, err
		}
		items = append(items, first_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
