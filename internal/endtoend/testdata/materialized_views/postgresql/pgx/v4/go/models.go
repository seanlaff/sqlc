// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package querytest

import (
	"database/sql"
)

type Author struct {
	ID     int64
	Name   string
	Bio    sql.NullString
	Gender sql.NullInt32
}

type AuthorsName struct {
	Name string
}
