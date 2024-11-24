package utils

import (
	"database/sql"
	"time"
)

// Converts a string to sql.NullString
func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

// Converts an sql.NullString to a Go string
// Returns empty string if Valid is false
func FromNullString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func ToNullInt32(i int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: i,
		Valid: true,
	}
}

// Converts time.Time into sql.NullTime
func ToNullTime(time time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  time,
		Valid: !time.IsZero(),
	}
}

// Converts an sql.NullTime to a time.Time
// Returns empty time.Time if Valid is false
func FromNullTime(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{}
}

// Returns the value of sql.NullString if Valid, or default value from param if not
func NullStringOrDefault(ns sql.NullString, defaultValue string) string {
	if ns.Valid {
		return ns.String
	}
	return defaultValue
}
