module api-oaxaca-com

go 1.13

require (
	api-oaxaca-com/packages/db v0.0.0
	api-oaxaca-com/packages/server v0.0.0
	api-oaxaca-com/resources/menu v0.0.0-00010101000000-000000000000
	api-oaxaca-com/resources/restaurant v0.0.0
)

replace (
	api-oaxaca-com/packages/db => ./packages/db
	api-oaxaca-com/packages/server => ./packages/server
	api-oaxaca-com/resources/menu => ./resources/menu
	api-oaxaca-com/resources/restaurant => ./resources/restaurant
	api-oaxaca-com/resources/staff => ./resources/staff
)
