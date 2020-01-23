package restaurant

// type restaurant represents a restaurant record in the database, and is used
// by the ORM to create the restaurants table. A unique group constraint is set
// on the name of the restaurant and its location - so that no two restaurants
// with the same name can exist at the same location.
type restaurant struct {
  RestaurantID int64 `pg:",pk"`
  Name string `pg:",notnull,unique:restaurant"`
  Location string `pg:",notnull,unique:restaurant"`
  TelephoneNumber string `pg:",notnull"`
}
