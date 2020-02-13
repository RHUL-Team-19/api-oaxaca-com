package authentication

import (
  "api-oaxaca-com/packages/db"
)

// type dbWrapper wraps a database connection and provides higher-level
// operations/abstractions on the database.
type dbWrapper struct {
  db *db.DB
}

// GetUser takes an id and fetches the user record in the database with that id.
func (w *dbWrapper) GetStaff(id int64) (*staff, error) {
  // construct empty record
  r := new(staff)

  // run query and handle error
  if err := w.db.Conn.
    Model(r).
    Where("staff_id = ?", id).
    Select(); err != nil {
      return nil, err
  } else {
    return r, nil
  }
}
