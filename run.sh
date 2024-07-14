#! /bin/bash
git pull

if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

docker-compose stop
docker-compose rm -f
docker-compose pull
docker-compose up
