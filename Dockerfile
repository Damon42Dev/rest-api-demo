# The final image is a lightweight image that contains only your app and its dependencies
FROM golang:1.21-bullseye AS builder
WORKDIR /rest-api-demo

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-docker-multistage

# final stage
FROM alpine:latest
COPY --from=builder /go-docker-multistage /go-docker-multistage
# Copy the .env file to the final image
COPY --from=builder /rest-api-demo/.env /.env
EXPOSE 8080
ENTRYPOINT [ "/go-docker-multistage" ]
