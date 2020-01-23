# resources
The API is split into multiple resources. A resource is an entity within the
system and database. A resource exposes HTTP GET/POST/PUT/PATCH/DELETE endpoints
and facilitates CRUD operations, as well as authorisation and input validation.
A resource must be passed to and mounted by api-oaxaca-com/packages/server.New
the server defines a resource interface which all resources must conform to.
