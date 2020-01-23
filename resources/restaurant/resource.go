package restaurant

// type Resource wraps a database connection and implements the
// api-oaxaca-com/packages/server.resource interface.
type Resource struct {
  db *dbWrapper
}
