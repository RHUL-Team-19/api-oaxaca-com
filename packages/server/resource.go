package server

import (
  "github.com/gin-gonic/gin"
)

// type resource defines an interface for server resources. A resource typically
// an entity with the system/database and allows HTTP CRUD operations on that
// entity.
type resource interface {
  AttachRoutes(e *gin.Engine)
  CreateTables() error
  CreateDefaultRecords() error
}
