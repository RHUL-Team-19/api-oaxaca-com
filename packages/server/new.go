package server

import (
  "github.com/gin-gonic/gin"
)

// New is a variadic function and takes any number of values of type resource.
// resource is an interface type. A gin engine is created and each resource's
// routes are attached to the engine. Each resource's database tables and
// default records are created.
func New(resources ...resource) (*Server, error) {
  // create gin engine
  r := gin.Default()

  for i := 0; i < len(resources); i++ {
    // attach resource's routes to engine
    resources[i].AttachRoutes(r)

    // create resource's table
    if err := resources[i].CreateTables(); err != nil {
      return nil, err
    }

    // insert default records into resource's tables
    if err := resources[i].CreateDefaultRecords(); err != nil {
      return nil, err
    }
  }

  // return pointer to Server type that wraps gin engine
  return &Server{ engine: r }, nil
}
