// Code generated by ent, DO NOT EDIT.

package entcinema

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the cinema type in the database.
	Label = "cinema"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldNumRow holds the string denoting the num_row field in the database.
	FieldNumRow = "num_row"
	// FieldNumColumn holds the string denoting the num_column field in the database.
	FieldNumColumn = "num_column"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldMinDistance holds the string denoting the min_distance field in the database.
	FieldMinDistance = "min_distance"
	// EdgeScreenings holds the string denoting the screenings edge name in mutations.
	EdgeScreenings = "screenings"
	// Table holds the table name of the cinema in the database.
	Table = "cinemas"
	// ScreeningsTable is the table that holds the screenings relation/edge.
	ScreeningsTable = "screenings"
	// ScreeningsInverseTable is the table name for the Screening entity.
	// It exists in this package in order to avoid circular dependency with the "screening" package.
	ScreeningsInverseTable = "screenings"
	// ScreeningsColumn is the table column denoting the screenings relation/edge.
	ScreeningsColumn = "cinema_screenings"
)

// Columns holds all SQL columns for cinema fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldNumRow,
	FieldNumColumn,
	FieldName,
	FieldMinDistance,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NumRowValidator is a validator for the "num_row" field. It is called by the builders before save.
	NumRowValidator func(uint32) error
	// NumColumnValidator is a validator for the "num_column" field. It is called by the builders before save.
	NumColumnValidator func(uint32) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// MinDistanceValidator is a validator for the "min_distance" field. It is called by the builders before save.
	MinDistanceValidator func(uint32) error
)

// OrderOption defines the ordering options for the Cinema queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByNumRow orders the results by the num_row field.
func ByNumRow(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumRow, opts...).ToFunc()
}

// ByNumColumn orders the results by the num_column field.
func ByNumColumn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumColumn, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByMinDistance orders the results by the min_distance field.
func ByMinDistance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMinDistance, opts...).ToFunc()
}

// ByScreeningsCount orders the results by screenings count.
func ByScreeningsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newScreeningsStep(), opts...)
	}
}

// ByScreenings orders the results by screenings terms.
func ByScreenings(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newScreeningsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newScreeningsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ScreeningsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ScreeningsTable, ScreeningsColumn),
	)
}
