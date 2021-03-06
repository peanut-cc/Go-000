// Code generated by entc, DO NOT EDIT.

package article

import (
	"time"
)

const (
	// Label holds the string label denoting the article type in the database.
	Label = "article"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldPublished holds the string denoting the published field in the database.
	FieldPublished = "published"
	// FieldCreatedTime holds the string denoting the created_time field in the database.
	FieldCreatedTime = "created_time"

	// Table holds the table name of the article in the database.
	Table = "articles"
)

// Columns holds all SQL columns for article fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldBody,
	FieldPublished,
	FieldCreatedTime,
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
	// DefaultCreatedTime holds the default value on creation for the created_time field.
	DefaultCreatedTime func() time.Time
)
