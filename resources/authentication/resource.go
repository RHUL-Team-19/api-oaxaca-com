package authentication

import (
  "github.com/gin-gonic/gin"
)

// type Resource wraps a database connection and implements the
// api-oaxaca-com/packages/server.resource interface.
type Resource struct {
  db *dbWrapper
}

// AttachRoutes mounts requests handlers to the gin engine.
func (m *Resource) AttachRoutes(e *gin.Engine) {
  // URL group
  g := e.Group("/authentication")

  // attach POST handler
  g.POST("/", m.postHandler)
}

// CreateTables does nothing here, and simply exists to satisfy the resource
// interface.
func (m *Resource) CreateTables() error {
  return nil
}

// CreateDefaultRecords does nothing here, and simply exists to satisfy the
// resource interface.
func (m *Resource) CreateDefaultRecords() error {
  return nil
}
