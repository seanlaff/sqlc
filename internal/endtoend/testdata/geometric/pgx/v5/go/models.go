// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type TestTable struct {
	VBoxNull     pgtype.Box
	VCircleNull  pgtype.Circle
	VLineNull    pgtype.Line
	VLsegNull    pgtype.Lseg
	VPathNull    pgtype.Path
	VPointNull   pgtype.Point
	VPolygonNull pgtype.Polygon
	VBox         pgtype.Box
	VCircle      pgtype.Circle
	VLine        pgtype.Line
	VLseg        pgtype.Lseg
	VPath        pgtype.Path
	VPoint       pgtype.Point
	VPolygon     pgtype.Polygon
}
