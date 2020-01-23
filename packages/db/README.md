# db
Package DB contains methods used by all resources. DB contains type DB which
wraps a Postgres database connection and provides high-level operations
including CreateTablesFor and DropTablesFor. Typically, each resource will wrap
type DB with its own database wrapper that provides more high-level database
methods that are specific to the service's needs.
