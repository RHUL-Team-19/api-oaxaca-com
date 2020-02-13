module api-oaxaca-com/resources/authentication

go 1.13

require (
	api-oaxaca-com/packages/auth v0.0.0
	api-oaxaca-com/packages/db v0.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
)

replace (
	api-oaxaca-com/packages/auth => ../../packages/auth
	api-oaxaca-com/packages/db => ../../packages/db
)
