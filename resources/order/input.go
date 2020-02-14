package order

// type createInput contains fields sent via a HTTP POST request to /orders
// which are used to create a new order record in the database.
type createInput struct {
  StaffID int64 `json:"staff_id"`
  TableID int64 `json:"table_id"`
  MealIDs []int64 `json:"meal_ids"`
}

// IsValid ensures the fields provided in createInput match a set of rules. If
// validation fails, a reason is returned - which should be sent to the client.
func (i *createInput) IsValid() (bool, string) {
  return true, ""
}

// type updateInput contains fields sent via a HTTP PUT/PATCH request to
// /orders/{id} which are used to fully update or partially update an
// order record in the database.
type updateInput struct {
  StaffID *int64 `json:"staff_id"`
  TableID *int64 `json:"table_id"`
  MealIDs *[]int64 `json:"meal_ids"`
}

// IsValid ensures the fields provided in updateInput match a set of rules. If
// validation fails, a reason is returned - which should be sent to the client.
func (i *updateInput) IsValid() (bool, string) {
  return true, ""
}
