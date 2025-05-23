// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cinema/pkg/ent/movie"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Movie is the model entity for the Movie schema.
type Movie struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration uint64 `json:"duration,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MovieQuery when eager-loading is set.
	Edges        MovieEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MovieEdges holds the relations/edges for other nodes in the graph.
type MovieEdges struct {
	// Screenings holds the value of the screenings edge.
	Screenings []*Screening `json:"screenings,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ScreeningsOrErr returns the Screenings value or an error if the edge
// was not loaded in eager-loading.
func (e MovieEdges) ScreeningsOrErr() ([]*Screening, error) {
	if e.loadedTypes[0] {
		return e.Screenings, nil
	}
	return nil, &NotLoadedError{edge: "screenings"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Movie) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case movie.FieldID, movie.FieldDuration:
			values[i] = new(sql.NullInt64)
		case movie.FieldTitle:
			values[i] = new(sql.NullString)
		case movie.FieldCreatedAt, movie.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Movie fields.
func (m *Movie) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case movie.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int64(value.Int64)
		case movie.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case movie.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case movie.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				m.Title = value.String
			}
		case movie.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				m.Duration = uint64(value.Int64)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Movie.
// This includes values selected through modifiers, order, etc.
func (m *Movie) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryScreenings queries the "screenings" edge of the Movie entity.
func (m *Movie) QueryScreenings() *ScreeningQuery {
	return NewMovieClient(m.config).QueryScreenings(m)
}

// Update returns a builder for updating this Movie.
// Note that you need to call Movie.Unwrap() before calling this method if this Movie
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Movie) Update() *MovieUpdateOne {
	return NewMovieClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Movie entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Movie) Unwrap() *Movie {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Movie is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Movie) String() string {
	var builder strings.Builder
	builder.WriteString("Movie(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(m.Title)
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", m.Duration))
	builder.WriteByte(')')
	return builder.String()
}

// Movies is a parsable slice of Movie.
type Movies []*Movie
