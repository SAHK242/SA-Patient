// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"patient/ent/treatdetail"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// TreatDetail is the model entity for the TreatDetail schema.
type TreatDetail struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TreatDetail) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case treatdetail.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TreatDetail fields.
func (td *TreatDetail) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case treatdetail.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			td.ID = int(value.Int64)
		default:
			td.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TreatDetail.
// This includes values selected through modifiers, order, etc.
func (td *TreatDetail) Value(name string) (ent.Value, error) {
	return td.selectValues.Get(name)
}

// Update returns a builder for updating this TreatDetail.
// Note that you need to call TreatDetail.Unwrap() before calling this method if this TreatDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (td *TreatDetail) Update() *TreatDetailUpdateOne {
	return NewTreatDetailClient(td.config).UpdateOne(td)
}

// Unwrap unwraps the TreatDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (td *TreatDetail) Unwrap() *TreatDetail {
	_tx, ok := td.config.driver.(*txDriver)
	if !ok {
		panic("ent: TreatDetail is not a transactional entity")
	}
	td.config.driver = _tx.drv
	return td
}

// String implements the fmt.Stringer.
func (td *TreatDetail) String() string {
	var builder strings.Builder
	builder.WriteString("TreatDetail(")
	builder.WriteString(fmt.Sprintf("id=%v", td.ID))
	builder.WriteByte(')')
	return builder.String()
}

// TreatDetails is a parsable slice of TreatDetail.
type TreatDetails []*TreatDetail
