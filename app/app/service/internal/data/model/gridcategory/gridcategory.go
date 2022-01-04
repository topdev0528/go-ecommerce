// Code generated by entc, DO NOT EDIT.

package gridcategory

import (
	"time"
)

const (
	// Label holds the string label denoting the gridcategory type in the database.
	Label = "grid_category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldImg holds the string denoting the img field in the database.
	FieldImg = "img"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCategoryID holds the string denoting the category_id field in the database.
	FieldCategoryID = "category_id"
	// FieldRootCategoryID holds the string denoting the root_category_id field in the database.
	FieldRootCategoryID = "root_category_id"
	// Table holds the table name of the gridcategory in the database.
	Table = "grid_category"
)

// Columns holds all SQL columns for gridcategory fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDeleteTime,
	FieldTitle,
	FieldImg,
	FieldName,
	FieldCategoryID,
	FieldRootCategoryID,
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)