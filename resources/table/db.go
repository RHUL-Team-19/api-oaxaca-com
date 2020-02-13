package table

import (
	"api-oaxaca-com/packages/db"
)

// type dbWrapper wraps a database connection and provides higher-level
// operations/abstractions on the database..
type dbWrapper struct {
	db *db.DB
}

// CreateTable takes a pointer to an instance of a createInput struct and
// then creates a new table model - which is then inserted into the
// database.
func (w *dbWrapper) CreateTable(i *createInput) error {
  // construct record
  r := table{
    RestaurantID: i.RestaurantID,
    TableNumber: i.TableNumber,
    NumberOfChairs: i.NumberOfChairs,
  }

  // insert record and handle error
  if _, err := w.db.Conn.
    Model(&r).
    Insert(); err != nil {
    return err
  } else {
    return nil
  }
}

// GetTable takes an id and fetches the table record in the database
// with that id.
func (w *dbWrapper) GetTable(id int64) (*table, error) {
	// construct empty record
	r := new(table)

	// run query and handle error
	if err := w.db.Conn.
		Model(r).
		Where("table_id = ?", id).
		Select(); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

// GetAllTables returns a slice containing all table records in the
// database.
func (w *dbWrapper) GetAllTables() ([]table, error) {
	// define space for result set
	var tables []table

	// run query and handle error
	if err := w.db.Conn.
		Model(&tables).
		Select(); err != nil {
		return nil, err
	} else {
		return tables, nil
	}
}

// GetAllTables takes a restaurant ID and returns a slice containing all table
// records in the database that belong to that restaurant.
func (w *dbWrapper) GetAllTablesForRestaurantID(id int64) ([]table, error) {
  // define space for result set
  var tables []table

  // run query and handle error
  if err := w.db.Conn.
    Model(&tables).
    Where("restaurant_id = ?", id).
    Select(); err != nil {
    return nil, err
  } else {
    return tables, nil
  }
}

// ...
func (w *dbWrapper) UpdateTable(id int64, i*updateInput) error {
  // construct empty record
  r := table{
    TableID: id,
  }

  // assign changed fields
  if i.RestaurantID != nil {
    r.RestaurantID = *i.RestaurantID
  }
  if i.TableNumber != nil {
    r.TableNumber = *i.TableNumber
  }
  if i.NumberOfChairs != nil {
    r.NumberOfChairs = *i.NumberOfChairs
  }

  // run query and handle error
  if _, err := w.db.Conn.
    Model(&r).
    WherePK().
    UpdateNotZero(); err != nil {
    return err
  } else {
    return nil
  }
}

// DeleteTable takes an id and deletes the record in the database with
// that id.
func (w *dbWrapper) DeleteTable(id int64) error {
	// construct record with ID only
	r := table{
		TableID: id,
	}

	// run query and handle error
	if _, err := w.db.Conn.
		Model(&r).
		WherePK().
		Delete(); err != nil {
		return err
	} else {
		return nil
	}
}
