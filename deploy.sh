docker build -t oaxaca-com/api-oaxaca-com . &&
  docker tag oaxaca-com/api-oaxaca-com registry.wsantos.net/oaxaca-com/api-oaxaca-com &&
  docker push registry.wsantos.net/oaxaca-com/api-oaxaca-com &&
  echo "done :)"
