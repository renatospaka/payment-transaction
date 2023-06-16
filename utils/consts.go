package utils

import "time"

const (
	ISO_FORMAT string = "2023-06-15T19:07:06.000"
)

var (
	NULL_DATE     time.Time = time.Time{}
	NULL_DATE_NIL string    = time.Time{}.String()
	NULL_DATE_STR string    = "null"
)
