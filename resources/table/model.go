package table

// type table represents a table record in the database, and is used by the
// ORM to create the table table.
type table struct {
  TableID int64 `pg:",pk" json:"table_id"`
  RestaurantID int64 `pg:",notnull,use_zero,unique:table" json:"restaurant_id"`
  TableNumber int `pg:",notnull,use_zero,unique:table" json:"table_number"`
  NumberOfChairs int `pg:",notnull,use_zero" json:"number_of_chairs"`
}
