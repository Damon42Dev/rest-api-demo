#How to run the project
1. to use the actual credentials for .env file (!! Do NOT push it to remote repo)
2. run "docker-compose up --build" in terminal
3. need to drop the database before every re-build (Will improve later on)

# Docker

### build docker image
docker build --rm -t go-docker-multistage:beta .

### run rest-api-demo project in docker container
docker-compose up --build

### stop container
docker-compose down

### to see docker images
docker image ls

### to see the running docker info
docker ps

## Portals
### portal of running project in docker container:
http://localhost:8080

### check host machine ip address
ifconfig | grep inet