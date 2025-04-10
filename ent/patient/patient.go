// Code generated by ent, DO NOT EDIT.

package patient

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the patient type in the database.
	Label = "patient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldDateOfBirth holds the string denoting the date_of_birth field in the database.
	FieldDateOfBirth = "date_of_birth"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// EdgeMedicalHistory holds the string denoting the medical_history edge name in mutations.
	EdgeMedicalHistory = "medical_history"
	// Table holds the table name of the patient in the database.
	Table = "patient"
	// MedicalHistoryTable is the table that holds the medical_history relation/edge.
	MedicalHistoryTable = "medical_histories"
	// MedicalHistoryInverseTable is the table name for the MedicalHistories entity.
	// It exists in this package in order to avoid circular dependency with the "medicalhistories" package.
	MedicalHistoryInverseTable = "medical_histories"
	// MedicalHistoryColumn is the table column denoting the medical_history relation/edge.
	MedicalHistoryColumn = "patient_id"
)

// Columns holds all SQL columns for patient fields.
var Columns = []string{
	FieldID,
	FieldPhone,
	FieldFirstName,
	FieldLastName,
	FieldGender,
	FieldAddress,
	FieldDateOfBirth,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
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
	// GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	GenderValidator func(int32) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Patient queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByDateOfBirth orders the results by the date_of_birth field.
func ByDateOfBirth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateOfBirth, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByMedicalHistoryCount orders the results by medical_history count.
func ByMedicalHistoryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMedicalHistoryStep(), opts...)
	}
}

// ByMedicalHistory orders the results by medical_history terms.
func ByMedicalHistory(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMedicalHistoryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMedicalHistoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MedicalHistoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MedicalHistoryTable, MedicalHistoryColumn),
	)
}
