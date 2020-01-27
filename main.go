package main

import (
  "api-oaxaca-com/packages/db"
  "api-oaxaca-com/packages/server"
  "api-oaxaca-com/resources/menu"
  "api-oaxaca-com/resources/restaurant"
)

const (
  // HTTP server config
  ServerHost = "0.0.0.0"
  ServerPort = "8080"

  // databse config
  DBHost = "127.0.0.1"
  DBPort = "5432"
  DBUser = "oaxaca"
  DBPass = "oaxaca"
  DBName = "oaxaca"
)

func main() {
  // connect to database
  pg := db.New(DBHost, DBPort, DBUser, DBPass, DBName)

  // create new server and provide resources - each resource receives a pointer
  // to the database connection
  s, err := server.New(
    menu.New(pg),
    restaurant.New(pg),
  )
  if err != nil {
    panic(err)
  }

  // start the server - blocks forever unless an error is returned
  if err := s.Start(ServerHost, ServerPort); err != nil {
    panic(err)
  }
}
