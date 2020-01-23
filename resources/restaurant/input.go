package restaurant

// type createInput contains fields sent via a HTTP POST request to /restaurants
// which are used to create a new restaurant record in the database.
type createInput struct {
  Name string `json:"name"`
  Location string `json:"location"`
  TelephoneNumber string `json:"telephone_number"`
}

// type updateInput contains fields sent via a HTTP PUT/PATCH request to
// /restaurants/{id} which are used to fully update or partially update a
// restaurant record in the database.
type updateInput struct {
  Name *string `json:"name"`
  Location *string `json:"location"`
  TelephoneNumber *string `json:"telephone_number"`
}
