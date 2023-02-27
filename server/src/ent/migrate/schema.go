// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TimeRecordsColumns holds the columns for the "time_records" table.
	TimeRecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_time_records", Type: field.TypeInt, Nullable: true},
	}
	// TimeRecordsTable holds the schema information for the "time_records" table.
	TimeRecordsTable = &schema.Table{
		Name:       "time_records",
		Columns:    TimeRecordsColumns,
		PrimaryKey: []*schema.Column{TimeRecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "time_records_users_timeRecords",
				Columns:    []*schema.Column{TimeRecordsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "manavis_id", Type: field.TypeInt, Unique: true},
		{Name: "grade", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TimeRecordsTable,
		UsersTable,
	}
)

func init() {
	TimeRecordsTable.ForeignKeys[0].RefTable = UsersTable
}
