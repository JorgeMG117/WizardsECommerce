# Step 1: Use an official Go image to build the Go app
FROM golang:1.20-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

RUN apk add --no-cache build-base sqlite-dev
ENV CGO_ENABLED=1


# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Build the Go app
RUN go build -o server cmd/server/main.go

RUN ls -l /app

# Step 2: Create a smaller image for the final build
FROM alpine:latest

# Install SQLite
#RUN apk add --no-cache sqlite
RUN apk add --no-cache sqlite-libs

# Set the working directory
WORKDIR /app

ENV GO_ENV=prod

RUN mkdir -p /app/static
RUN mkdir -p /app/views

# Copy the built Go binary from the build stage
COPY --from=build /app/server/main /app/server/main


COPY server.crt .
COPY server.key .

COPY static /app/static
COPY views /app/views

# Expose the port your app runs on (adjust if needed)
EXPOSE 8080

# Command to run the executable
CMD ["/app/server/main"]
