#How to run the project
1. to use the actual credentials for .env file (!! Do NOT push it to remote repo)
2. run "docker-compose up --build" in terminal

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


## Run Test cases
### Goto controller folder
cd /Users/damonwang/go/src/example/rest-api-demo/src/controllers
### Run all test cases
go test -v
### Run all a specific test case
go test -v -run TEST_FUNCTION_NAME