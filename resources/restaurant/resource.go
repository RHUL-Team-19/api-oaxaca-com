package restaurant

// type Resource wraps a database connection and implements the
// api-oaxaca-com/packages/server.resource interface.
type Resource struct {
  Conn *dbWrapper
}

// AttachRoutes mounts requests handlers to the gin engine.
func (m *Resource) AttachRoutes(e *gin.Engine) {
  // URL group
  g := e.Group("/restaurants")

  // attach POST handler
  g.POST("/", postHandler)

  // attach GET handlers
  g.GET("/", getAllHandler)
  g.GET("/:id", getOneHandler)

  // attach PUT handlers
  g.PUT("/:id", putHandler)

  // attach PATCH handlers
  g.PATCH("/:id", patchHandler)

  // attach DELETE handelrs
  g.DELETE("/:id", deleteHandler)
}

// CreateTables takes all models and creates tables for them in the database.
func (m *Resource) CreateTables() error {
  return m.db.Conn.CreateTablesFor(
    new(restaurant),
  )
}

// CreateDefaultRecords does nothing here, and simply exists to satisfy the
// resource interface.
func (m *Resource) CreateDefaultRecords() error {
  return nil
}
