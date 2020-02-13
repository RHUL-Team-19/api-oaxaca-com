package table

import (
  "testing"
)

// TestCreateTable tests the CreateTable db method by creating a fake
// table record and then trying to retrieve it from the database. If these two
// steps complete without error, we assume the test passes.
func TestCreateTable(t *testing.T) {

  // ...

}

// TestGetTable tests the GetTable db method by creating a fake table
// record and then retrieving it by its id. If both of these steps complete
// without error, we assume the test passes.
func TestGetTable(t *testing.T) {

  // ...

}

// TestGetAllTables tests the GetAllTables db method by creating multiple
// fake table records and then retrieving all of those records. If both of these
// steps complete without error, we assume the test passes.
func TestGetAllTables(t *testing.T) {

  // ...

}

// TestGetAllTablesForRestaurantID tests the GetAllTablesForRestaurantID db
// method by creating multiple fake table records and then retrieving all of
// those records. If both of these steps complete without error, we assume the
// test passes.
func TestGetAllTablesForRestaurantID(t *testing.T) {

  // ...

}

// TestUpdateTable tests the UpdateTable db method by creating a fake
// table record and then updating one of its fields. The updated record is
// then retrieved from the database, and the updated field is checked. If these
// steps complete without error, we assume the test passes.
func TestUpdateTable(t *testing.T) {

  // ...

}

// TestDeleteTable tests the DeleteTable db method by creating a fake
// record and then deleting it. Then tries to re-fetch the deleted record to
// make sure it doesn't exist. If these steps complete without error, we assume
// the test passes.
func TestDeleteTable(t *testing.T) {

  // ...

}
