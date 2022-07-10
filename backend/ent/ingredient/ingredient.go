// Code generated by entc, DO NOT EDIT.

package ingredient

const (
	// Label holds the string label denoting the ingredient type in the database.
	Label = "ingredient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the ingredient in the database.
	Table = "ingredients"
)

// Columns holds all SQL columns for ingredient fields.
var Columns = []string{
	FieldID,
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