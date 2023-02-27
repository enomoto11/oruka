// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldManavisID holds the string denoting the manavis_id field in the database.
	FieldManavisID = "manavis_id"
	// FieldGrade holds the string denoting the grade field in the database.
	FieldGrade = "grade"
	// EdgeTimeRecords holds the string denoting the timerecords edge name in mutations.
	EdgeTimeRecords = "timeRecords"
	// Table holds the table name of the user in the database.
	Table = "users"
	// TimeRecordsTable is the table that holds the timeRecords relation/edge.
	TimeRecordsTable = "time_records"
	// TimeRecordsInverseTable is the table name for the TimeRecord entity.
	// It exists in this package in order to avoid circular dependency with the "timerecord" package.
	TimeRecordsInverseTable = "time_records"
	// TimeRecordsColumn is the table column denoting the timeRecords relation/edge.
	TimeRecordsColumn = "user_time_records"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldFirstName,
	FieldLastName,
	FieldManavisID,
	FieldGrade,
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
	// FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// ManavisIDValidator is a validator for the "manavis_id" field. It is called by the builders before save.
	ManavisIDValidator func(int) error
	// GradeValidator is a validator for the "grade" field. It is called by the builders before save.
	GradeValidator func(int) error
)
