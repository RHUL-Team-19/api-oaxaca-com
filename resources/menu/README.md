# menu
This resource defines a New function that takes a database and returns a pointer
to a new instance of Resource. This can then be passed to a module
api-oaxaca-com/packages/server.New function.

### models
```
menuMeal {
  meal_id: ID!
  name: String!
  price: Float!
  description: String!
  is_vegan: Boolean!
  is_vegetarian: Boolean!
  does_contain_egg: Boolean!
  does_contain_soy: Boolean!
  does_contain_fish: Boolean!
  does_contain_lactose: Boolean!
  does_contain_wheat: Boolean!
  does_contain_nuts: Boolean!
  does_contain_gluten: Boolean!
  does_contain_dairy: Boolean!
  image_url: String!
}
```
This package defines a menu meal model - which represents a meal (which is on
the restaurant's menu) record within the database.

### inputs and input validation
```
create_input {
  name: String!
  price: Float!
  description: String!
  is_vegan: Boolean!
  is_vegetarian: Boolean!
  does_contain_egg: Boolean!
  does_contain_soy: Boolean!
  does_contain_fish: Boolean!
  does_contain_lactose: Boolean!
  does_contain_wheat: Boolean!
  does_contain_nuts: Boolean!
  does_contain_gluten: Boolean!
  does_contain_dairy: Boolean!
  image_url: String!
}

update_input {
  name: String
  price: Float
  description: String
  is_vegan: Boolean
  is_vegetarian: Boolean
  does_contain_egg: Boolean
  does_contain_soy: Boolean
  does_contain_fish: Boolean
  does_contain_lactose: Boolean
  does_contain_wheat: Boolean
  does_contain_nuts: Boolean
  does_contain_gluten: Boolean
  does_contain_dairy: Boolean
  image_url: String
}
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
