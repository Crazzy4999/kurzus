package util

import "database/sql"

func NullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}
