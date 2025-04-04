// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"patient/ent/invoice"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Invoice is the model entity for the Invoice schema.
type Invoice struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Invoice) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case invoice.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Invoice fields.
func (i *Invoice) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case invoice.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Invoice.
// This includes values selected through modifiers, order, etc.
func (i *Invoice) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// Update returns a builder for updating this Invoice.
// Note that you need to call Invoice.Unwrap() before calling this method if this Invoice
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Invoice) Update() *InvoiceUpdateOne {
	return NewInvoiceClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Invoice entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Invoice) Unwrap() *Invoice {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Invoice is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Invoice) String() string {
	var builder strings.Builder
	builder.WriteString("Invoice(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Invoices is a parsable slice of Invoice.
type Invoices []*Invoice
