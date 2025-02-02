# Docker

### build docker image
docker build --rm -t go-docker-multistage:beta .

### run rest-api-demo project in docker container
docker run -d -p 8080:8081 --name rest-api-demo-app go-docker-multistage:beta

### to see docker images
docker image ls

### to see the running docker info
docker ps

# Notes
* Docker image needs to be rebuilt once any change is made. 

## Portals
### portal of running project locally:
http://localhost:8081
### portal of running project in docker container:
http://localhost:8080
