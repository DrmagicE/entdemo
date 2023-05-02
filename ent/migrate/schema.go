// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CarsColumns holds the columns for the "cars" table.
	CarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// CarsTable holds the schema information for the "cars" table.
	CarsTable = &schema.Table{
		Name:       "cars",
		Columns:    CarsColumns,
		PrimaryKey: []*schema.Column{CarsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cars_users_cars",
				Columns:    []*schema.Column{CarsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// GroupUsersBindingColumns holds the columns for the "group_users_binding" table.
	GroupUsersBindingColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeInt},
	}
	// GroupUsersBindingTable holds the schema information for the "group_users_binding" table.
	GroupUsersBindingTable = &schema.Table{
		Name:       "group_users_binding",
		Columns:    GroupUsersBindingColumns,
		PrimaryKey: []*schema.Column{GroupUsersBindingColumns[0], GroupUsersBindingColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_users_binding_user_id",
				Columns:    []*schema.Column{GroupUsersBindingColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_users_binding_group_id",
				Columns:    []*schema.Column{GroupUsersBindingColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CarsTable,
		GroupsTable,
		UsersTable,
		GroupUsersBindingTable,
	}
)

func init() {
	CarsTable.ForeignKeys[0].RefTable = UsersTable
	GroupUsersBindingTable.ForeignKeys[0].RefTable = GroupsTable
	GroupUsersBindingTable.ForeignKeys[1].RefTable = UsersTable
}
