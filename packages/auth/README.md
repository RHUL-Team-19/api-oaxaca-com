# auth (authentication and authorisation)
Package authentication contains methods for hashing passwords, comparing hashes,
and generating JSON web tokens (JWTs). Also contains a middleware used by
request handlers that are restricted to requests that contain a JWT with
specific roles.
