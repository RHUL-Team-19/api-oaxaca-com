package staff

// type createInput contains fields sent via a HTTP POST request to /staffs
// which are used to create a new staff record in the database.
type createInput struct {
  RestaurantID int64 `json:"restaurant_id"`
  FullName string `json:"full_name"`
  UserName string `json:"user_name"`
  PasswordHash string `json:"password_hash"`
  HasPassedTraining bool `json:"has_passed_training"`
}

// IsValid ensures the fields provided in createInput match a set of rules. If
// validation fails, a reason is returned - which should be sent to the client.
func (i *createInput) IsValid() (bool, string) {
  if i.RestaurantID < 0 {
    return false, "RestaurantID cannot be less than 0"
  }

  if len(i.FullName) < 2 || len(i.FullName) > 40 {
    return false, "Length of full name must be between 2 and 40"
  }

  if len(i.UserName) < 1 || len(i.UserName) > 20 {
    return false, "Length of user name must be between 1 and 20"
  }

  if len(i.PasswordHash) < 1 || len(i.PasswordHash) > 16 {
    return false, "Length of password hash must be between 1 and 16"
  }

  return true, ""
}

// type updateInput contains fields sent via a HTTP PUT/PATCH request to
// /staffs/{id} which are used to fully update or partially update a
// staff record in the database.
type updateInput struct {
  RestaurantID int64 `json:"restaurant_id"`
  FullName string `json:"full_name"`
  UserName string `json:"user_name"`
  PasswordHash string `json:"password_hash"`
  HasPassedTraining bool `json:"has_passed_training"`
}

// IsValid ensures the fields provided in updateInput match a set of rules. If
// validation fails, a reason is returned - which should be sent to the client.
func (i *updateInput) IsValid() (bool, string) {
  if i.RestaurantID < 0 {
    return false, "RestaurantID cannot be less than 0"
  }

  if len(i.FullName) < 2 || len(i.FullName) > 40 {
    return false, "Length of full name must be between 2 and 40"
  }

  if len(i.UserName) < 1 || len(i.UserName) > 20 {
    return false, "Length of user name must be between 1 and 20"
  }

  if len(i.PasswordHash) < 1 || len(i.PasswordHash) > 16 {
    return false, "Length of password hash must be between 1 and 16"
  }

  return true, ""
}
