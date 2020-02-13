package staff

// type staff represents a staff record in the database, and is used by the
// ORM to create the staff table.
type staff struct {
  StaffID int64 `pg:",pk" json:"staff_id"`
  RestaurantID int64 `pg:",fk" json:"restaurant_id"`
  FullName string `pg:",notnull" json:"full_name"`
  Score int `pg:",notnull,use_zero" json:"score"`
  PasswordHash []byte `pg:",notnull" json:"password_hash"`
  HasPassedTraining bool `pg:",notnull,use_zero" json:"has_passed_training"`
  Role string `pg:",notnull" json:"role"`
}
