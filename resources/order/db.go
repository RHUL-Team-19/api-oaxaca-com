package order

import (
	"api-oaxaca-com/packages/db"
)

// type dbWrapper wraps a database connection and provides higher-level
// operations/abstractions on the database..
type dbWrapper struct {
	db *db.DB
}
