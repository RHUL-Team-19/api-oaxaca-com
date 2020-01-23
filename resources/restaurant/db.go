package restaurant

import (
  "api-oaxaca-com/packages/db"
)

// type dbWrapper wraps a database connection and provides higher-level
// operations/abstractions on the database.
type dbWrapper struct {
  Conn *db.DB
}

// CreateRestaurant takes a pointer to an instance of a createInput struct and
// then creates a new restaurant model - which is then inserted into the
// database.
func (w *dbWrapper) CreateRestaurant(i *createInput) error {
  // construct record
  r := restaurant{
    Name: i.Name,
    Location: i.Location,
    TelephoneNumber: i.TelephoneNumber,
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

// GetRestaurant takes an id and fetches the restaurant record in the database
// with that id.
func (w *dbWrapper) GetRestaurant(id int64) (*restaurant, error) {
  // construct record with ID only
  r := restaurant{
    RestaurantID: id,
  }

  // run query and handle error
  if _, err := w.db.Conn.
    Model(&r).
    Select(); err != nil {
      return nil, err
  } else {
    return &r, nil
  }
}

// GetAllRestaurants returns a slice containing all restaurant records in the
// database.
func (w *dbWrapper) GetAllRestaurants() ([]restaurant, error) {
  // define space for result set
  var restaurants []restaurant

  // run query and handlr error
  if _, err := w.db.Conn.
    Model(&restaurants).
    Select(); err != nil {
      return nil, err
  } else {
    return restaurants, nil
  }
}

// UpdateRestaurant takes an id and a pointer to an instance of an updateInput
// struct and then fetches the restaurant record in the database with the
// corresponding id - and updates its fields.
func (w *dbWrapper) UpdateRestaurant(id int64, i *updateInput) error {

  // ...

  return nil
}

// DeleteRestaurant takes an id and deletes the record in the database with
// that id.
func (w *dbWrapper) DeleteRestaurant(id int64) error {

  // ...

  return nil
}
