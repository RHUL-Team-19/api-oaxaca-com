# api-oaxaca-com
The backend REST API for Oaxaca

The API is split into multiple resources. A resource is an entity within the
system - and thus one or more tables in the database. Example resources include:
- restaurant
- menu
- staff
The API allows CRUD (create, retrieve, update, delete) operations on these
resources.

### making requests and handling responses
All data in a GET response payload will be serialised as JSON. A request which
returns more than one object - via a list - will return null (instead of an
empty list) if the result set is empty.
A GET or DELETE to a specific resource record via its ID will return an error
message if no record with that ID exists, along with an appropriate HTTP status
code.

Data sent to the server via POST and PATCH - as an input - must be serialised as
JSON. All valid inputs are described below, and information regarding nullable
fields and input validation can be found in each resource's readme file.
If a field in an input fails a validation check, an error message (and
appropriate HTTP status code) containing an reason/explanation will be returned.

### authentication and authorisation
In order to access most of the API resource endpoints, one must first
authenticate by making a POST request containing user credentials to the
/authentication resource. A valid set of credentials will return a payload
containing a JSON web token (JWT) which must be sent in the "Authorization"
header in future requests. This will allow access to restricted resource
endpoints - which are defined below.
```
Authorization: bearer <JWT token string>
```

### access scope
A token has a role associated with it. Some restricted endpoints require a
specific role. E.g. in order to create a new restaurant, one must be of the
role "manager". Roles are defined below. If a user's role is changed, a new
token must be issued. Valid roles are:
- waiter
- kitchen
- manager
Some endpoints do not require a role - i.e. they are publicly available. This
is denoted with an "all" role.

### resource record ids
A resource record ID is a number in string form, and uniquely identifies a
specific record for a specific resource. Record IDs are usually used in a URL
when sending a GET or DELETE request to a specific record. E.g:
```
GET https://some-url.io/restaurants/12
DELETE https://some-url.io/restaurants/15
```
A record ID can also be sent via a create or update input when creating a
relation between two entities. For example, when creating a new member of staff:
```json
{
  "restaurant_id": "12"
}
```
Record IDs are always sent as strings, never integer values.

### error messages
Error messages are returned when:
- an input contains invalid field names
- an input fails a validation check
- an internal error occurs
- a GET or DELETE is requested on a resource record ID that is non-existent
and are of the form:
```json
{
  "error": "<message text>"
}
```

## /menu
 Path  | Method | Access scope | Input             | Possible responses                                      | Description
:------|:------:|:-------------|:------------------|:--------------------------------------------------------|------------------
 /     | GET    | All          | N/A               | 200 with JSON payload or 400 with error                 | Fetch all meals on the menu
 /{id} | GET    | All          | N/A               | 200 with JSON payload or null                           | Fetch a specific meal by its ID
 /     | POST   | Manager      | create input      | 201 with no payload or 400 with error or 500 with error | Add new meal to the menu
 /{id} | PATCH  | Manager      | update input      | 201 with no payload or 400 with error or 500 with error | Update a meal on the menu by its ID
 /{id} | DELETE | Manager      | N/A               | 201 with no payload or 400 with error or 500 with error | Remove a meal from the menu

### inputs
#### create input
 Attribute name       | Type    | Required
:---------------------|:--------|:--------:
 name                 | string  | yes
 price                | float   | yes
 description          | string  | yes
 is_vegan             | boolean | yes
 is_vegetarian        | boolean | yes
 does_contain_egg     | boolean | yes
 does_contain_soy     | boolean | yes
 does_contain_fish    | boolean | yes
 does_contain_nuts    | boolean | yes
 does_contain_wheat   | boolean | yes
 does_contain_dairy   | boolean | yes
 does_contain_gluten  | boolean | yes
 does_contain_lactose | boolean | yes
 image_url            | string  | yes

#### update input
 Attribute name       | Type    | Required
:---------------------|:--------|:--------:
 name                 | string  | no
 price                | float   | no
 description          | string  | no
 is_vegan             | boolean | no
 is_vegetarian        | boolean | no
 does_contain_egg     | boolean | no
 does_contain_soy     | boolean | no
 does_contain_fish    | boolean | no
 does_contain_nuts    | boolean | no
 does_contain_wheat   | boolean | no
 does_contain_dairy   | boolean | no
 does_contain_gluten  | boolean | no
 does_contain_lactose | boolean | no
 image_url            | string  | no


## /restaurant
 Path  | Method | Access scope | Input             | Possible responses                                      | Description
:------|:------:|:-------------|:------------------|:--------------------------------------------------------|------------------
 /     | GET    | Manager      | N/A               | 200 with JSON payload or 400 with error                 | Fetch all restaurants
 /{id} | GET    | Manager      | N/A               | 200 with JSON payload or null                           | Fetch a specific restaurant by its ID
 /     | POST   | Manager      | create input      | 201 with no payload or 400 with error or 500 with error | Create a new restaurant
 /{id} | PATCH  | Manager      | update input      | 201 with no payload or 400 with error or 500 with error | Update a restaurant by its ID
 /{id} | DELETE | Manager      | N/A               | 201 with no payload or 400 with error or 500 with error | Remove a restaurant

### inputs
#### create input
 Attribute name   | Type    | Required
:-----------------|:--------|:--------:
 name             | string  | yes
 location         | string  | yes
 telephone_number | string  | yes

#### update input
 Attribute name   | Type    | Required
:-----------------|:--------|:--------:
 name             | string  | yes
 location         | string  | yes
 telephone_number | string  | yes


## /staff
 Path  | Method | Access scope | Input             | Possible responses                                      | Description
:------|:------:|:-------------|:------------------|:--------------------------------------------------------|------------------
 /     | GET    | Manager      | N/A               | 200 with JSON payload or 400 with error                 | Fetch all members of staff
 /{id} | GET    | Waiter       | N/A               | 200 with JSON payload or null                           | Fetch a specific member of staff by their ID
 /     | POST   | Manager      | create input      | 201 with no payload or 400 with error or 500 with error | Create a new member of staff
 /{id} | PATCH  | Manager      | update input      | 201 with no payload or 400 with error or 500 with error | Update a member of staff by their ID
 /{id} | DELETE | Manager      | N/A               | 201 with no payload or 400 with error or 500 with error | Delete a member of staff

### inputs
#### create input
 Attribute name | Type | Required
:---------------|:-----|:--------:

#### update input
 Attribute name | Type | Required
:---------------|:-----|:--------:


## /authentication
 Path  | Method | Access scope | Input                | Possible responses                                      | Description
:------|:------:|:-------------|:---------------------|:--------------------------------------------------------|:---------------------------------
 /     | POST   | All          | authentication input |                                                         | Request a new token

### inputs
#### authentication input
 Attribute name | Type | Required
:---------------|:-----|:--------:
