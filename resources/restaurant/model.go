package restaurant

// type restaurant represents a restaurant record in the database, and is used
// by the ORM to create the restaurants table.
type restaurant struct {
  RestaurantID int64 `pg:",pk"`
  Name string `pg:",notnull,unique:restaurant"`
  Location string `pg:",notnull,unique:restaurant"`
  TelephoneNumber string `pg:",notnull"`
}
