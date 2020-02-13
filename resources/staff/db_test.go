package staff

import (
  "testing"
)

// TestCreateStaff tests the CreateStaff db method by creating a fake
// staff record and then trying to retrieve it from the database. If these two
// steps complete without error, we assume the test passes.
func TestCreateStaff(t *testing.T) {

  // ...

}

// TestGetStaff tests the GetStaff db method by creating a fake staff
// record and then retrieving it by its id. If both of these steps complete
// without error, we assume the test passes.
func TestGetStaff(t *testing.T) {

  // ...

}

// TestGetAllStaff tests the GetAllStaff db method by creating multiple
// fake staff records and then retrieving all of those records. If both of these
// steps complete without error, we assume the test passes.
func TestGetAllStaff(t *testing.T) {

  // ...

}

// TestUpdateStaff tests the UpdateStaff db method by creating a fake
// staff record and then updating one of its fields. The updated record is
// then retrieved from the database, and the updated field is checked. If these
// steps complete without error, we assume the test passes.
func TestUpdateStaff(t *testing.T) {

  // ...

}

// TestDeleteStaff tests the DeleteStaff db method by creating a fake
// record and then deleting it. Then tries to re-fetch the deleted record to
// make sure it doesn't exist. If these steps complete without error, we assume
// the test passes.
func TestDeleteStaff(t *testing.T) {

  // ...

}
