FROM golang:1.22-alpine AS builder

LABEL maintainer="Snehalkumar Mahale <svmahale1991@gmail.com>"

# Create a directory to store build files
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Build the Go app
RUN go build -o apiserver

#############################################################################
FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder /build/apiserver .

# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]