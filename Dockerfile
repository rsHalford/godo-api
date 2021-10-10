# syntax=docker/dockerfile:1

ARG BUILDER_IMAGE=golang:bullseye
ARG DISTROLESS_IMAGE=gcr.io/distroless/base-debian11

# Build

# Base image for the go application
FROM ${BUILDER_IMAGE} AS builder

RUN update-ca-certificates

# Set current working directory inside container
WORKDIR /app

# Copy dependency files to the PWD
COPY go.mod .

ENV GO111MODULE=on

# Download Go module dependencies
RUN go mod download
RUN go mod verify

COPY . .

# Build GoDo API binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o main .


# Deploy

FROM ${DISTROLESS_IMAGE}

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port 8080 of the container
EXPOSE 8080

# Run the executable
ENTRYPOINT ["./main"]
