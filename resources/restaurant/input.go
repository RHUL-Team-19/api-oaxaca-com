package restaurant

// type createInput contains fields sent via a HTTP POST request to /restaurants
// which are used to create a new restaurant record in the database.
type createInput struct {
  Name string `json:"name"`
  Location string `json:"location"`
  TelephoneNumber string `json:"telephone_number"`
}

// IsValid ensures the fields provided in createInput match a set of rules. If
// validation fails, a reason is returned - which should be sent to the client.
func (i *createInput) IsValid() (bool, string) {
  if len(i.Name) < 3 {
    return false, "Name must be at least 3 characters in length"
  }

  if len(i.Location) < 5 {
    return false, "Location must be at least 5 characters in length"
  }

  if len(i.TelephoneNumber) < 5 {
    return false, "Telephone number must be at least 5 characters in length"
  }

  return true, ""
}

// type updateInput contains fields sent via a HTTP PUT/PATCH request to
// /restaurants/{id} which are used to fully update or partially update a
// restaurant record in the database.
type updateInput struct {
  Name *string `json:"name"`
  Location *string `json:"location"`
  TelephoneNumber *string `json:"telephone_number"`
}

// IsValid ensures the fields provided in updateInput match a set of rules. If
// validation fails, a reason is returned - which should be sent to the client.
func (i *updateInput) IsValid() (bool, string) {
  if i.Name != nil && len(*i.Name) < 3 {
    return false, "Name must be at least 3 characters in length"
  }

  if i.Location != nil && len(*i.Location) < 5 {
    return false, "Location must be at least 5 characters in length"
  }

  if i.TelephoneNumber != nil && len(*i.TelephoneNumber) < 5 {
    return false, "Telephone number must be at least 5 characters in length"
  }

  return true, ""
}
