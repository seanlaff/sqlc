// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type TestTable struct {
	VCidNull pgtype.Uint32
	VOidNull pgtype.Uint32
	VTidNull pgtype.TID
	VXidNull pgtype.Uint32
	VCid     pgtype.Uint32
	VOid     pgtype.Uint32
	VTid     pgtype.TID
	VXid     pgtype.Uint32
}
