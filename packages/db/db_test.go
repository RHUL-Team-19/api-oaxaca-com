package db

import (
  "testing"
)

const (
  // databse config for tests
  testDBHost = "127.0.0.1"
  testDBPort = "5432"
  testDBUser = "oaxaca"
  testDBPass = "oaxaca"
  testDBName = "oaxaca"
)

// type testTableA defines an example table containing a primary key field,
// a non-nullable name field, and a nullable address field.
type testTableA struct {
  ID int64 `pg:",pk"`
  Name string `pg:",notnull"`
  Address string
}

// type testTableB defines an example table containing a primary key, a
// non-nullable height field, and a nullable weight field.
type testTableB struct {
  ID int64 `pg:",pk"`
  Height float32 `pg:",notnull,use_zero"`
  Weight float32
}

// TestCreateTablesFor creates a new database client and attempts to create two
// test tables. Then inserts test records into the tables. Then queries the
// tables. If these three steps complete without error, we assume the
// CreateTablesFor passes.
func TestCreateTablesFor(t *testing.T) {
  // connect
  pg := New(testDBHost, testDBPort, testDBUser, testDBPass, testDBName)

  // create tables
  if err := pg.CreateTablesFor(
    new(testTableA),
    new(testTableB),
  ); err != nil {
    t.Error("Cannot create tables: ", err.Error())
  }

  // insert example records for test table A
  for _, r := range []testTableA{
    testTableA{
      Name: "test name",
    },
    testTableA{
      Name: "another test name",
      Address: "some test address",
    },
    testTableA{
      Name: "yet another test name",
      Address: "another test address",
    },
  } {
    if _, err := pg.Conn.
      Model(&r).
      Insert(); err != nil {
      t.Error("Cannot create test A record: ", err.Error())
    }
  }

  // insert example records for test table B
  for _, r := range []testTableB{
    testTableB{},
    testTableB{
      Height: 1.23,
    },
    testTableB{
      Height: 2.34,
      Weight: 5.76,
    },
  } {
    if _, err := pg.Conn.
      Model(&r).
      Insert(); err != nil {
      t.Error("Cannot create test B record: ", err.Error())
    }
  }

  // make sure test table A can be queried
  var resultsA []testTableA
  if err := pg.Conn.
    Model(&resultsA).
    Select(); err != nil {
    t.Error("Cannot query test A table: ", err.Error())
  }

  // make sure test table B can be queried
  var resultsB []testTableB
  if err := pg.Conn.
    Model(&resultsB).
    Select(); err != nil {
    t.Error("Cannot query test B table: ", err.Error())
  }
}

// TestDropTablesFor creates a new database client and attempts to drop two
// test tables. Then queries the tables. If the queries fail, we assume
// TestDropTablesFor passed since there must be no tables to query.
func TestDropTablesFor(t *testing.T) {
  // connect
  pg := New(testDBHost, testDBPort, testDBUser, testDBPass, testDBName)

  // drop tables
  if err := pg.DropTablesFor(
    new(testTableA),
    new(testTableB),
  ); err != nil {
    t.Error("Cannot drop tables: ", err.Error())
  }

  // make sure test table A cannot be queried
  var resultsA []testTableA
  if err := pg.Conn.
    Model(&resultsA).
    Select(); err == nil {
    t.Error("Can query test A table after it was dropped")
  }

  // make sure test table B cannot be queried
  var resultsB []testTableB
  if err := pg.Conn.
    Model(&resultsB).
    Select(); err == nil {
    t.Error("Can query test B table after it was dropped")
  }
}
