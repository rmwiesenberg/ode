#! /bin/bash
git checkout main
git pull

docker-compose down
docker-compose pull

if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi
docker-compose up --build
