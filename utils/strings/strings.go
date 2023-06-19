package strings

import (
	"database/sql"
	"log"
)

// When database/sql (postgres) returns a date/time value, it may be null
// So it is necessary to verify when the value is valid (time is a date/time)
// or invalid (time is zero date/tim) before converting to the proper format
func FormatNullString(nullString sql.NullString) string {
	log.Printf("FormatNullString - nullString.Valid: %t, nullString.String: %v\n", nullString.Valid, nullString.String)
	if nullString.Valid {
		return nullString.String
	} 
	return ""
}
