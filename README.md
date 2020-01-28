# api-oaxaca-com
The backend REST API for Oaxaca

## /menu
 Path  | Method | Access scope | Input             | Response      | Description
:------|:------:|:-------------|:------------------|:--------------|------------------
 /     | GET    | All          | N/A               |               | Fetch all meals on the menu
 /{id} | GET    | All          | N/A               |               | Fetch a specific meal by its ID
 /     | POST   | Manager      | create input      |               | Add new meal to the menu
 /{id} | PATCH  | Manager      | update input      |               | Update a meal on the menu by its ID
 /{id} | DELETE | Manager      | N/A               |               | Remove a meal from the menu

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
 does_contain_wheat   | boolean | yes
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
 does_contain_wheat   | boolean | no
 does_contain_gluten  | boolean | no
 does_contain_lactose | boolean | no
 image_url            | string  | no


## /restaurant
 Path  | Method | Access scope | Input             | Response      | Description
:------|:------:|:-------------|:------------------|:--------------|------------------
 /     | GET    | Manager      | N/A               |               | Fetch all restaurants
 /{id} | GET    | Manager      | N/A               |               | Fetch a specific restaurant by its ID
 /     | POST   | Manager      | create input      |               | Create a new restaurant
 /{id} | PATCH  | Manager      | update input      |               | Update a restaurant by its ID
 /{id} | DELETE | Manager      | N/A               |               | Remove a restaurant

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
 Path  | Method | Access scope | Input             | Response      | Description
:------|:------:|:-------------|:------------------|:--------------|------------------
 /     | GET    | Manager      | N/A               |               | Fetch all members of staff
 /{id} | GET    | Waiter       | N/A               |               | Fetch a specific member of staff by their ID
 /     | POST   | Manager      | create input      |               | Create a new member of staff
 /{id} | PATCH  | Manager      | update input      |               | Update a member of staff by their ID
 /{id} | DELETE | Manager      | N/A               |               | Delete a member of staff

### inputs
#### create input
 Attribute name | Type | Required
:---------------|:-----|:--------:

#### update input
 Attribute name | Type | Required
:---------------|:-----|:--------:


## /authentication
 Path  | Method | Access scope | Input                | Response      | Description
:------|:------:|:-------------|:---------------------|:--------------|:---------------------------------
 /     | POST   | All          | authentication input |               | Request a new token

### inputs
#### authentication input
 Attribute name | Type | Required
:---------------|:-----|:--------:
