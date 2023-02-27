// Code generated by ent, DO NOT EDIT.

package timerecord

import (
	"time"
)

const (
	// Label holds the string label denoting the timerecord type in the database.
	Label = "time_record"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTimeKeeper holds the string denoting the timekeeper edge name in mutations.
	EdgeTimeKeeper = "timeKeeper"
	// Table holds the table name of the timerecord in the database.
	Table = "time_records"
	// TimeKeeperTable is the table that holds the timeKeeper relation/edge.
	TimeKeeperTable = "time_records"
	// TimeKeeperInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	TimeKeeperInverseTable = "users"
	// TimeKeeperColumn is the table column denoting the timeKeeper relation/edge.
	TimeKeeperColumn = "user_time_records"
)

// Columns holds all SQL columns for timerecord fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "time_records"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_time_records",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
)
