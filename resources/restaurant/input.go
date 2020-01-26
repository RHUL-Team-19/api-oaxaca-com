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
  if len(i.Name) < 3 || len(i.Name) > 20 {
    return false, "Length of name must be between 3 and 20"
  }

  if len(i.Location) < 5 || len(i.Location) > 20 {
    return false, "Length of location must be between 5 and 20"
  }

  if len(i.TelephoneNumber) < 5 || len(i.TelephoneNumber) > 20 {
    return false, "Length of telephone number must be between 5 and 20"
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
  if i.Name != nil && (len(*i.Name) < 3 || len(*i.Name) > 20) {
    return false, "Length of name must be between 3 and 20"
  }

  if i.Location != nil && (len(*i.Location) < 5 || len(*i.Location) > 20) {
    return false, "Length of location must be between 5 and 20"
  }

  if i.TelephoneNumber != nil && (len(*i.TelephoneNumber) < 5 ||
    len(*i.TelephoneNumber) > 20) {
      return false, "Length of telephone number must be between 5 and 20"
  }

  return true, ""
}
