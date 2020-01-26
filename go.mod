module api-oaxaca-com

go 1.13

require (
	api-oaxaca-com/packages/db v0.0.0
	api-oaxaca-com/packages/server v0.0.0
	api-oaxaca-com/resources/restaurant v0.0.0
)

replace (
	api-oaxaca-com/packages/db => ./packages/db
	api-oaxaca-com/packages/server => ./packages/server
	api-oaxaca-com/resources/restaurant => ./resources/restaurant
)
