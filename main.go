package main

import (
  "github.com/0xc0392b/envy"
  "api-oaxaca-com/packages/db"
  "api-oaxaca-com/packages/server"
  "api-oaxaca-com/resources/menu"
  "api-oaxaca-com/resources/restaurant"
  "api-oaxaca-com/resources/staff"
)

var (
  // HTTP server config
  ServerHost = envy.GetEnvAsStr("SERVER_HOST", "0.0.0.0")
  ServerPort = envy.GetEnvAsStr("SERVER_PORT", "8080")

  // databse config
  DBHost = envy.GetEnvAsStr("DB_HOST", "127.0.0.1")
  DBPort = envy.GetEnvAsStr("DB_PORT", "5432")
  DBUser = envy.GetEnvAsStr("DB_USER", "oaxaca")
  DBPass = envy.GetEnvAsStr("DB_PASS", "oaxaca")
  DBName = envy.GetEnvAsStr("DB_NAME", "oaxaca")
)

func main() {
  // connect to database
  pg := db.New(DBHost, DBPort, DBUser, DBPass, DBName)

  // create new server and provide resources - each resource receives a pointer
  // to the database connection
  s, err := server.New(
    menu.New(pg),
    restaurant.New(pg),
    staff.New(pg),
  )
  if err != nil {
    panic(err)
  }

  // start the server - blocks forever unless an error is returned
  if err := s.Start(ServerHost, ServerPort); err != nil {
    panic(err)
  }
}
