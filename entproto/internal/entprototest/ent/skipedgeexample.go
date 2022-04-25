// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/contrib/entproto/internal/entprototest/ent/skipedgeexample"
	"entgo.io/contrib/entproto/internal/entprototest/ent/user"
	"entgo.io/ent/dialect/sql"
)

// SkipEdgeExample is the model entity for the SkipEdgeExample schema.
type SkipEdgeExample struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SkipEdgeExampleQuery when eager-loading is set.
	Edges          SkipEdgeExampleEdges `json:"edges"`
	user_skip_edge *int
}

// SkipEdgeExampleEdges holds the relations/edges for other nodes in the graph.
type SkipEdgeExampleEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SkipEdgeExampleEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SkipEdgeExample) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case skipedgeexample.FieldID:
			values[i] = new(sql.NullInt64)
		case skipedgeexample.ForeignKeys[0]: // user_skip_edge
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SkipEdgeExample", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SkipEdgeExample fields.
func (see *SkipEdgeExample) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case skipedgeexample.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			see.ID = int(value.Int64)
		case skipedgeexample.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_skip_edge", value)
			} else if value.Valid {
				see.user_skip_edge = new(int)
				*see.user_skip_edge = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the SkipEdgeExample entity.
func (see *SkipEdgeExample) QueryUser() *UserQuery {
	return (&SkipEdgeExampleClient{config: see.config}).QueryUser(see)
}

// Update returns a builder for updating this SkipEdgeExample.
// Note that you need to call SkipEdgeExample.Unwrap() before calling this method if this SkipEdgeExample
// was returned from a transaction, and the transaction was committed or rolled back.
func (see *SkipEdgeExample) Update() *SkipEdgeExampleUpdateOne {
	return (&SkipEdgeExampleClient{config: see.config}).UpdateOne(see)
}

// Unwrap unwraps the SkipEdgeExample entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (see *SkipEdgeExample) Unwrap() *SkipEdgeExample {
	tx, ok := see.config.driver.(*txDriver)
	if !ok {
		panic("ent: SkipEdgeExample is not a transactional entity")
	}
	see.config.driver = tx.drv
	return see
}

// String implements the fmt.Stringer.
func (see *SkipEdgeExample) String() string {
	var builder strings.Builder
	builder.WriteString("SkipEdgeExample(")
	builder.WriteString(fmt.Sprintf("id=%v", see.ID))
	builder.WriteByte(')')
	return builder.String()
}

// SkipEdgeExamples is a parsable slice of SkipEdgeExample.
type SkipEdgeExamples []*SkipEdgeExample

func (see SkipEdgeExamples) config(cfg config) {
	for _i := range see {
		see[_i].config = cfg
	}
}