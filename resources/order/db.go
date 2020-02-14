package order

import (
	"time"
	"api-oaxaca-com/packages/db"
)

// type dbWrapper wraps a database connection and provides higher-level
// operations/abstractions on the database..
type dbWrapper struct {
	db *db.DB
}

// CreateOrder takes a pointer to an instance of a createInput struct and
// then creates a new order model - which is then inserted into the
// database.
func (w *dbWrapper) CreateOrder(i *createInput) error {
  // construct record
  r := order{
		StaffID: i.StaffID,
		TableID: i.TableID,
		DateTimeCreated: time.Now(),
  }

  // insert record and handle error
  if _, err := w.db.Conn.
    Model(&r).
    Insert(); err != nil {
    return err
  }

	// create an order meal record for each meal ID
	for _, mealID := range i.MealIDs {
		if err := w.CreateMealOrder(&orderMeal{
			OrderID: r.OrderID,
			MealID: mealID,
		}); err != nil {
			return err
		}
	}

	// no error
	return nil
}

// CreateMealOrder takes a pointer to an orderMeal record and inserts it into
// the database.
func (w *dbWrapper) CreateMealOrder(r *orderMeal) error {
	// insert record and handle error
  if _, err := w.db.Conn.
    Model(r).
    Insert(); err != nil {
    return err
  } else {
    return nil
  }
}

// GetOrder takes an id and fetches the order record in the database
// with that id.
func (w *dbWrapper) GetOrder(id int64) (*order, error) {
	// construct empty record
	r := new(order)

	// run query and handle error
	if err := w.db.Conn.
		Model(r).
		Where("order_id = ?", id).
		Select(); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

// GetAllOrders returns a slice containing all order records in the
// database.
func (w *dbWrapper) GetAllOrders() ([]order, error) {
	// define space for result set
	var orders []order

	// run query and handle error
	if err := w.db.Conn.
		Model(&orders).
		Select(); err != nil {
		return nil, err
	} else {
		return orders, nil
	}
}

// GetAllOrdersForRestaurantID takes a restaurant ID and returns a slice
// containing all order records in the database that belong to that restaurant.
func (w *dbWrapper) GetAllOrdersForRestaurantID(id int64) ([]order, error) {
  // define space for result set
  var orders []order

  // run query and handle error
  if err := w.db.Conn.
    Model(&orders).
		Join("NATURAL JOIN tables").
    Where("restaurant_id = ?", id).
    Select(); err != nil {
    return nil, err
  } else {
    return orders, nil
  }
}

// GetAllOrdersForTableID takes a table ID and returns a slice containing all
// order records in the database that belong to that table.
func (w *dbWrapper) GetAllOrdersForTableID(id int64) ([]order, error) {
  // define space for result set
  var orders []order

  // run query and handle error
  if err := w.db.Conn.
    Model(&orders).
    Where("table_id = ?", id).
    Select(); err != nil {
    return nil, err
  } else {
    return orders, nil
  }
}

// GetMealsForOrderID takes an order ID and fetches and retuens all meal orders
// for that order from the database.
func (w *dbWrapper) GetMealsForOrderID(id int64) ([]orderMeal, error) {
	// define space for result set
  var meals []orderMeal

  // run query and handle error
  if err := w.db.Conn.
    Model(&meals).
    Where("order_id = ?", id).
    Select(); err != nil {
    return nil, err
  } else {
    return meals, nil
  }
}

// UpdateOrder takes an id and a pointer to an instance of an updateInput
// struct and then fetches the order record in the database with the
// corresponding id - and updates its fields.
func (w *dbWrapper) UpdateOrder(id int64, i*updateInput) error {
  // construct empty record
  r := order{
    OrderID: id,
  }

  // assign changed fields
  if i.StaffID != nil {
    r.StaffID = *i.StaffID
  }
	if i.TableID != nil {
    r.TableID = *i.TableID
  }
	if i.SatisfactionRating != nil {
		r.SatisfactionRating = *i.SatisfactionRating
	}
	if i.MealIDs != nil {
		// delete all existing order meals
		if err := w.DeleteMealOrdersForOrderID(id); err != nil {
			return err
		}

		// create an order meal record for each meal ID
		for _, mealID := range *i.MealIDs {
			if err := w.CreateMealOrder(&orderMeal{
				OrderID: r.OrderID,
				MealID: mealID,
			}); err != nil {
				return err
			}
		}
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

// DeleteOrder takes an id and deletes the record in the database with
// that id.
func (w *dbWrapper) DeleteOrder(id int64) error {
	// construct record with ID only
	r := order{
		OrderID: id,
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

// DeleteMealOrdersForOrderID takes an order ID and deletes all meal order
// records for it.
func (w *dbWrapper) DeleteMealOrdersForOrderID(id int64) error {
	if _, err := w.db.Conn.
		Model(new(orderMeal)).
		Where("order_id = ?", id).
		Delete(); err != nil {
		return err
	} else {
		return nil
	}
}
