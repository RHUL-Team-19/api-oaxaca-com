# api-oaxaca-com
The backend REST API for Oaxaca

## /menu
 Path  | Method | Access scope | Input             | Description
:------|:------:|:-------------|:------------------|:---------------------------------
 /     | GET    | All          | N/A               | Fetch all meals on the menu
 /{id} | GET    | All          | N/A               | Fetch a specific meal by its ID
 /     | POST   | Manager      | create input      | Add new meal to the menu
 /{id} | PATCH  | Manager      | update input      | Update a meal on the menu by its ID
 /{id} | DELETE | Manager      | N/A               | Remove a meal from the menu

### inputs
#### create input
 Attribute name | Type | Required
:---------------|:-----|:--------:

#### update input
 Attribute name | Type | Required
:---------------|:-----|:--------:


## /restaurant


## /staff


## /authentication
