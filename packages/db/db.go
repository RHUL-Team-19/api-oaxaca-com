package db

import (
  "github.com/go-pg/pg/v9"
  "github.com/go-pg/pg/v9/orm"
)

// DB wraps a Postgres server connection
type DB struct {
  Conn *pg.DB
}

// CreateTableFor takes a list of model struct and creates the necessary table
// in the database if it doesn't already exist.
func (d *DB) CreateTablesFor(models ...interface{}) error {
  for _, model := range models {
    if err := d.Conn.CreateTable(
      model,
      &orm.CreateTableOptions{
        IfNotExists: true,
      },
    ); err != nil {
      return err
    }
  }
  return nil
}

// DropTableFor takes a list of model struct and drops the table from the
// database if it exists. Drops also cascade.
func (d *DB) DropTablesFor(models ...interface{}) error {
  for _, model := range models {
    if err := d.Conn.DropTable(
      model,
      &orm.DropTableOptions{
        IfExists: true,
        Cascade: true,
      },
    ); err != nil {
      return err
    }
  }
  return nil
}
