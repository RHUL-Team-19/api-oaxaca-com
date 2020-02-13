# authentication
This resource defines a New function that takes a database and returns a pointer
to a new instance of Resource. This can then be passed to a module
api-oaxaca-com/packages/server.New function.

### models
```
authResponse {
  token: String!
}
```
This package defines an auth response model - which wraps a JSON web token (JWT)
sent to the user if they provide valid credentials.

### inputs and input validation
```
authentication_input {
  userID: String!
  password: String!
}
```
When requesting a new JWT, a authentication_input must be provided.

### HTTP POST
```
POST /authentication
```
Send user ID and password and request a new JWT.
