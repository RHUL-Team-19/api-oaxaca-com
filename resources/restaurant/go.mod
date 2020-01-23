module api-oaxaca-com/resources/restaurant

go 1.13

require (
  api-oaxaca-com/packages/db v0.0.0
)

replace (
  api-oaxaca-com/packages/db => ../../packages/db
)
