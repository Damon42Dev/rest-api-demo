# # The final image is a lightweight image that contains only your app and its dependencies
FROM golang:1.19.2-bullseye AS builder
WORKDIR /rest-api-demo
COPY go.mod .
RUN go mod download
COPY *.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/go-docker-multistage

# final stage
FROM alpine:latest
COPY --from=builder /opt/go-docker-multistage /opt/go-docker-multistage
EXPOSE 8080
ENTRYPOINT [ "/opt/go-docker-multistage" ]


# Below is the Dockerfile for a test build
# # Specifies a parent image
# FROM golang:1.19.2-bullseye
 
# # Creates an app directory to hold your app’s source code
# WORKDIR /rest-api-demo
 
# # Copies everything from your root directory into /rest-api-demo
# COPY go.mod .
 
# # Installs Go dependencies
# RUN go mod download

# # Copies everything from your root directory into /rest-api-demo
# COPY main.go .
 
# # Builds your app with optional configuration
# RUN go build -o /godocker
 
# # Tells Docker which network port your container listens on
# EXPOSE 8080
 
# # Specifies the executable command that runs when the container starts
# CMD [ "/godocker" ]