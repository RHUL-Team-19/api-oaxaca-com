package order

import (
  "testing"
)

// TestCreateOrder tests the CreateOrder db method by creating a fake
// order record and then trying to retrieve it from the database. If these two
// steps complete without error, we assume the test passes.
func TestCreateOrder(t *testing.T) {

  // ...

}

// TestGetOrder tests the GetPrder db method by creating a fake order
// record and then retrieving it by its id. If both of these steps complete
// without error, we assume the test passes.
func TestGetOrder(t *testing.T) {

  // ...

}

// TestGetAllOrders tests the GetAllOrders db method by creating multiple
// fake order records and then retrieving all of those records. If both of these
// steps complete without error, we assume the test passes.
func TestGetAllOrders(t *testing.T) {

  // ...

}

// TestGetAllOrdersForRestaurantID tests the GetAllOrdersForRestaurantID db
// method by creating multiple fake order records and then retrieving all of
// those records. If both of these steps complete without error, we assume the
// test passes.
func TestGetAllOrdersForRestaurantID(t *testing.T) {

  // ...

}

// TestGetAllOrdersForTableID tests the GetAllOrdersForTableID db
// method by creating multiple fake order records and then retrieving all of
// those records. If both of these steps complete without error, we assume the
// test passes.
func TestGetAllOrdersForTableID(t *testing.T) {

  // ...

}

// TestUpdateOrder tests the UpdateOrder db method by creating a fake
// order record and then updating one of its fields. The updated record is
// then retrieved from the database, and the updated field is checked. If these
// steps complete without error, we assume the test passes.
func TestUpdateOrder(t *testing.T) {

  // ...

}

// TestDeleteOrder tests the DeleteOrder db method by creating a fake
// record and then deleting it. Then tries to re-fetch the deleted record to
// make sure it doesn't exist. If these steps complete without error, we assume
// the test passes.
func TestDeleteOrder(t *testing.T) {

  // ...

}
