module api-oaxaca-com

go 1.13

require (
	api-oaxaca-com/packages/db v0.0.0
	api-oaxaca-com/packages/server v0.0.0
	api-oaxaca-com/resources/menu v0.0.0
	api-oaxaca-com/resources/restaurant v0.0.0
	github.com/0xc0392b/envy v0.0.0-20200129005622-8213f89f5335
)

replace (
	api-oaxaca-com/packages/db => ./packages/db
	api-oaxaca-com/packages/server => ./packages/server
	api-oaxaca-com/resources/menu => ./resources/menu
	api-oaxaca-com/resources/restaurant => ./resources/restaurant
	api-oaxaca-com/resources/staff => ./resources/staff
)
