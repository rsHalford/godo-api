# syntax=docker/dockerfile:1

# Build

# Base image for the go application
FROM golang:1.17-bullseye AS build

# Set current working directory inside container
WORKDIR /app/goapi-server

# Copy dependency files to the PWD
COPY go.mod ./
COPY go.sum ./

# Download Go module dependencies
RUN go mod download

COPY *.go ./
COPY config/ ./
COPY model/ ./

# Build GoAPI binary and name it /server
RUN go build -o /goapi-server

# Deploy

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /goapi-server /goapi-server

# Expose port 8080 of the container
EXPOSE 8080

USER nonroot:nonroot

# Run the executable
ENTRYPOINT ["/goapi-server"]
