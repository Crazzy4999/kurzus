package request

import "database/sql"

type UserRequest struct {
	Address   sql.NullInt64 `json:"address"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
}
