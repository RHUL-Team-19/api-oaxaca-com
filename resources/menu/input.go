package menu

// type createInput contains fields sent via a HTTP POST request to /menu
// which are used to create a new meal record in the database.
type createInput struct {}

// IsValid ensures the fields provided in createInput match a set of rules. If
// validation fails, a reason is returned - which should be sent to the client.
func (i *createInput) IsValid() (bool, string) {

  // ...

  return true, ""
}

// type updateInput contains fields sent via a HTTP PUT/PATCH request to
// /menu/{id} which are used to fully update or partially update a menu record
// in the database.
type updateInput struct {}

// IsValid ensures the fields provided in updateInput match a set of rules. If
// validation fails, a reason is returned - which should be sent to the client.
func (i *updateInput) IsValid() (bool, string) {

  // ...

  return true, ""
}
