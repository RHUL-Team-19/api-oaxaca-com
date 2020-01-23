package server

import (
  "testing"
  "github.com/gin-gonic/gin"
)

// type testResourceA is a simple test resource that defines a path name.
type testResourceA struct {
  pathName string
}

// AttachRoutes attaches the resource's path to the engine, and defines an empty
// request handler.
func (r testResourceA) AttachRoutes(e *gin.Engine) {
  e.GET(r.pathName, func(c *gin.Context) {})
}

// CreateTables exists just to satisfy the resouce interface.
func (r testResourceA) CreateTables() error {
  return nil
}

// CreateDefaultRecords exists just to satisfy the resouce interface.
func (r testResourceA) CreateDefaultRecords() error {
  return nil
}

// type testResourceA is a simple test resource that defines a path name.
type testResourceB struct {
  pathName string
}

// AttachRoutes attaches the resource's path to the engine, and defines an empty
// request handler.
func (r testResourceB) AttachRoutes(e *gin.Engine) {
  e.GET(r.pathName, func(c *gin.Context) {})
}

// CreateTables exists just to satisfy the resouce interface.
func (r testResourceB) CreateTables() error {
  return nil
}

// CreateDefaultRecords exists just to satisfy the resouce interface.
func (r testResourceB) CreateDefaultRecords() error {
  return nil
}

// TestNew creates a new instance of server and then passes two test resources
// to it. Those resources are mounted to the server and are accessible via the
// paths they define. The resources are stumps - they don't actually do
// anything, and are just for testing purposes.
func TestNew(t *testing.T) {
  if _, err := New(
    &testResourceA{
      pathName: "/test-resource-A",
    },
    &testResourceB{
      pathName: "/test-resource-b",
    },
  ); err != nil {
    t.Error("Unable to create server and mount resources: ", err.Error())
  }
}
