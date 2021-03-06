# restaurant
This resource defines a New function that takes a database and returns a pointer
to a new instance of Resource. This can then be passed to a module
api-oaxaca-com/packages/server.New function.

### models
```
restaurant {
  restaurant_id: ID!
  name: String!
  telephone_number: String!
  location: String!
}
```
This package defines a restaurant model - which represents a restaurant record
within the database.

### inputs and input validation
```
create_input {
  name: String! // 3 <= length <= 20; non-nullable field
  telephone_number: String! // 5 <= length <= 20; non-nullable field
  location: String! // 5 <= length <= 20; non-nullable field
}

update_input {
  name: String // 3 <= length <= 20; nullable field
  telephone_number: String // 5 <= length <= 20; nullable field
  location: String // 5 <= length <= 20; nullable field
}
```
When POSTing to the server, a create_input must be provided. When updating an
existing record in the database, an update_input must be provided along with
the record's id.

### HTTP POST
```
POST /restaurants
```
Creates a new restaurant record.

### HTTP GET
```
GET /restaurants
GET /restaurants/{id}
```
Fetch all restaurant records or fetch one restaurant by ID.

### HTTP PATCH
```
PATCH /restaurants/{id}
```
Partial update for an existing restaurant record.

### HTTP DELETE
```
DELETE /restaurants/{id}
```
Delete a restaurant record by its ID.
