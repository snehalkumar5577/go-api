# Step 1: Build the Go application with CGO enabled
FROM golang:1.22-alpine AS builder

LABEL maintainer="Snehalkumar Mahale <svmahale1991@gmail.com>"

# Create a directory to store build files
WORKDIR /build

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o apiserver

#############################################################################
# Step 2: Create a minimal runtime image
FROM alpine:latest

# Copy the built binary from the builder stage
COPY --from=builder /build/apiserver .

# Command to run the application
ENTRYPOINT ["/apiserver"]
