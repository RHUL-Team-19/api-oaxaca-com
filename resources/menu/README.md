# menu
This resource defines a New function that takes a database and returns a pointer
to a new instance of Resource. This can then be passed to a module
api-oaxaca-com/packages/server.New function.

### models
```
menuMeal {
  meal_id: ID!
}
```
This package defines a menu meal model - which represents a meal (which is on
the restaurant's menu) record within the database.

### inputs and input validation
```
create_input {}

update_input {}
```
When POSTing to the server, a create_input must be provided. When updating an
existing record in the database, an update_input must be provided along with
the record's id.

### HTTP POST
```
POST /menu
```
Creates a new menu meal record.

### HTTP GET
```
GET /menu
GET /menu/{id}
```
Fetch all menu meal records or fetch one meal by ID.

### HTTP PATCH
```
PATCH /menu/{id}
```
Partial update for an existing menu meal record.

### HTTP DELETE
```
DELETE /meal/{id}
```
Delete a meal record by its ID.
