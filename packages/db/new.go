package db

import (
  "github.com/go-pg/pg/v9"
)

// New takes a Postgres server host, port, user, password, and database name
// string and returns a new DB.
func New(host, port, user, pass, name string) *DB {
  return &DB{
    Conn: pg.Connect(&pg.Options{
      Addr: host + ":" + port,
      User: user,
      Password: pass,
      Database: name,
    }),
  }
}
