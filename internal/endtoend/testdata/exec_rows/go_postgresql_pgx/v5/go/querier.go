// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package querytest

import (
	"context"
)

type Querier interface {
	DeleteBarByID(ctx context.Context, id int32) (int64, error)
}

var _ Querier = (*Queries)(nil)
