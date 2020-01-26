package menu

import (
  "api-oaxaca-com/packages/db"
)

// type dbWrapper wraps a database connection and provides higher-level
// operations/abstractions on the database.
type dbWrapper struct {
  db *db.DB
}

// CreateMenuMeal takes a pointer to an instance of a createInput struct and
// then creates a new meal model - which is then inserted into the database.
func (w *dbWrapper) CreateMenuMeal(i *createInput) error {
  // construct record
  r := menuMeal{}

  // insert record and handle error
  if _, err := w.db.Conn.
    Model(&r).
    Insert(); err != nil {
      return err
  } else {
    return nil
  }
}

// GetMenuMeal takes an id and fetches the meal record in the database with that
// id.
func (w *dbWrapper) GetMenuMeal(id int64) (*menuMeal, error) {
  // construct record with ID only
  r := menuMeal{
    MealID: id,
  }

  // run query and handle error
  if err := w.db.Conn.
    Model(&r).
    Select(); err != nil {
      return nil, err
  } else {
    return &r, nil
  }
}

// GetAllMenuMeals returns a slice containing all meal records in the database.
func (w *dbWrapper) GetAllMenuMeals() ([]menuMeal, error) {
  // define space for result set
  var meals []menuMeal

  // run query and handle error
  if err := w.db.Conn.
    Model(&meals).
    Select(); err != nil {
      return nil, err
  } else {
    return meals, nil
  }
}

// UpdateMenuMeal takes an id and a pointer to an instance of an updateInput
// struct and then fetches the meal record in the database with the
// corresponding id - and updates its fields.
func (w *dbWrapper) UpdateMenuMeal(id int64, i *updateInput) error {
  // construct empty record
  r := menuMeal{
    MealID: id,
  }

  // assign changed fields
  // ...

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

// DeleteMenuMeal takes an id and deletes the record in the database with
// that id.
func (w *dbWrapper) DeleteMenuMeal(id int64) error {
  // construct record with ID only
  r := menuMeal{
    MealID: id,
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
