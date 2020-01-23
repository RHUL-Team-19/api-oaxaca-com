package restaurant

import (
  "api-oaxaca-com/packages/db"
)

// type dbWrapper wraps a database connection and provides higher-level
// operations/abstractions on the database.
type dbWrapper struct {
  db *db.DB
}

// CreateRestaurant takes a pointer to an instance of a createInput struct and
// then creates a new restaurant model - which is then inserted into the
// database.
func (w *dbWrapper) CreateRestaurant(i *createInput) error {

  // ...

  return nil
}

// GetRestaurant takes an id and fetches the restaurant record in the database
// with that id.
func (w *dbWrapper) GetRestaurant(id int64) (*restaurant, error) {

  // ...

  return nil, nil
}

// GetAllRestaurants returns a slice containing all restaurant records in the
// database.
func (w *dbWrapper) GetAllRestaurants() ([]restaurant, error) {

  // ...

  return nil, nil
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
