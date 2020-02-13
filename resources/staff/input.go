package staff

// type createInput contains fields sent via a HTTP POST request to /staffs
// which are used to create a new staff record in the database.
type createInput struct {
  RestaurantID int64 `json:"restaurant_id"`
  FullName string `json:"full_name"`
  Score int `json:"score"`
  Password string `json:"password"`
  HasPassedTraining bool `json:"has_passed_training"`
  Role string `json:"role"`
}

// IsValid ensures the fields provided in createInput match a set of rules. If
// validation fails, a reason is returned - which should be sent to the client.
func (i *createInput) IsValid() (bool, string) {
  if len(i.FullName) < 2 || len(i.FullName) > 60 {
    return false, "Length of full name must be between 2 and 60"
  }

  if i.Score < 0 {
    return false, "Score cannot be less than 0"
  }

  if len(i.Password) < 6 || len(i.Password) > 32 {
    return false, "Length of password hash must be between 6 and 32"
  }

  if len(i.Role) > 20 {
    return false, "Length of role cannot be greater than 20"
  }

  return true, ""
}

// type updateInput contains fields sent via a HTTP PUT/PATCH request to
// /staffs/{id} which are used to fully update or partially update a
// staff record in the database.
type updateInput struct {
  RestaurantID *int64 `json:"restaurant_id"`
  FullName *string `json:"full_name"`
  Score *int `json:"score"`
  Password *string `json:"password"`
  HasPassedTraining *bool `json:"has_passed_training"`
  Role *string `json:"role"`
}

// IsValid ensures the fields provided in updateInput match a set of rules. If
// validation fails, a reason is returned - which should be sent to the client.
func (i *updateInput) IsValid() (bool, string) {
  if i.FullName != nil && (len(*i.FullName) < 2 || len(*i.FullName) > 60) {
    return false, "Length of full name must be between 2 and 60"
  }

  if i.Score != nil && *i.Score < 0 {
    return false, "Score cannot be less than 0"
  }

  if i.Password != nil && (len(*i.Password) < 6 || len(*i.Password) > 32) {
    return false, "Length of password hash must be between 6 and 32"
  }

  if i.Role != nil && len(*i.Role) > 20 {
    return false, "Length of role cannot be greater than 20"
  }

  return true, ""
}
