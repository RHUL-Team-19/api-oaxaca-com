package restaurant

import (
  "testing"
)

// TestCreateRestaurant tests the CreateRestaurant db method by creating a fake
// restaurant record and then trying to retrieve it from the database. If these
// two steps complete without error, we assume the test passes.
func TestCreateRestaurant(t *testing.T) {

  // ...

}

// TestGetRestaurant tests the GetRestaurant db method by creating a fake
// restaurant record and then retrieving it by its id. If both of these steps
// complete without error, we assume the test passes.
func TestGetRestaurant(t *testing.T) {

  // ...

}

// TestGetAllRestaurants tests the GetAllRestaurants db method by creating
// multiple fake restaurant records and then retrieving all of those records.
// If both of these steps complete without error, we assume the test passes.
func TestGetAllRestaurants(t *testing.T) {

  // ...

}

// TestUpdateRestaurant tests the UpdateRestaurant db method by creating a fake
// restaurant record and then updating one of its fields. The updated record is
// then retrieved from the database, and the updated field is checked. If these
// steps complete without error, we assume the test passes.
func TestUpdateRestaurant(t *testing.T) {

  // ...

}

// TestDeleteRestaurant tests the DeleteRestaurant db method by creating a fake
// record and then deleting it. Then tries to re-fetch the deleted record to
// make sure it doesn't exist. If these steps complete without error, we assume
// the test passes.
func TestDeleteRestaurant(t *testing.T) {

  // ...

}
