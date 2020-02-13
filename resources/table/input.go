package table

// type createInput contains fields sent via a HTTP POST request to /tables
// which are used to create a new table record in the database.
type createInput struct {
  RestaurantID int64 `json:"restaurant_id"`
  TableNumber int `json:"table_number"`
  NumberOfChairs int `json:"number_of_chairs"`
}

// IsValid ensures the fields provided in createInput match a set of rules. If
// validation fails, a reason is returned - which should be sent to the client.
func (i *createInput) IsValid() (bool, string) {
  if i.TableNumber < 0 {
    return false, "Table number cannot be less than 0"
  }

  if i.NumberOfChairs < 1 || i.NumberOfChairs > 30 {
    return false, "Number of chairs must be between 1 and 30"
  }

  return true, ""
}

// type updateInput contains fields sent via a HTTP PUT/PATCH request to
// /tables/{id} which are used to fully update or partially update a
// table record in the database.
type updateInput struct {
  RestaurantID *int64 `json:"restaurant_id"`
  TableNumber *int `json:"table_number"`
  NumberOfChairs *int `json:"number_of_chairs"`
}

// IsValid ensures the fields provided in updateInput match a set of rules. If
// validation fails, a reason is returned - which should be sent to the client.
func (i *updateInput) IsValid() (bool, string) {
  if i.TableNumber != nil && *i.TableNumber < 0 {
    return false, "Table number cannot be less than 0"
  }

  if i.NumberOfChairs != nil && (*i.NumberOfChairs < 1 ||
    *i.NumberOfChairs > 30) {
      return false, "Number of chairs must be between 1 and 30"
  }

  return true, ""
}
