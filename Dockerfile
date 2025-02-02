# The final image is a lightweight image that contains only your app and its dependencies
FROM golang:1.19.2-bullseye AS builder
WORKDIR /rest-api-demo
# Ensure both go.mod and go.sum are copied
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
# Copy the .env file
COPY .env ./
# List files to verify .env is copied
RUN ls -la 
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/go-docker-multistage

# final stage
FROM alpine:latest
COPY --from=builder /opt/go-docker-multistage /opt/go-docker-multistage
# Copy the .env file to the final image
COPY --from=builder /rest-api-demo/.env /opt/.env
# List files to verify .env is copied in the final image
RUN ls -la /opt
EXPOSE 8080
ENTRYPOINT [ "/opt/go-docker-multistage" ]
