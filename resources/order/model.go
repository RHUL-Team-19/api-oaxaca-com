package order

import (
  "time"
)

// type order represents an order record in the database, and is used by the
// ORM to create the order table.
type order struct {
  OrderID int64 `pg:",pk" json:"order_id"`
  StaffID int64 `pg:",notnull,use_zero" json:"staff_id"`
  TableID int64 `pg:",notnull,use_zero" json:"table_id"`
  SatisfactionRating int `json:"satisfaction_rating"`
  DateTimeCreated time.Time `pg:",notnull" json:"date_time_created"`
}

// type orderMeal represents an meal that belongs to an order record in the
// database, and is used by the ORM to create the order meal table.
type orderMeal struct {
  OrderMealID int64 `pg:",pk" json:"order_meal_id"`
  OrderID int64 `pg:",notnull,use_zero" json:"order_id"`
  MealID int64 `pg:",notnull,use_zero" json:"meal_id"`
  DateTimePrepared time.Time `json:"date_time_prepared"`
}
