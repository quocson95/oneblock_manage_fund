# Use the official Golang image to create a build artifact.
# This image has the Go compiler pre-installed.
FROM golang:1.23.0 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -ldflags="-s -w" -o oneblock_manage_fund_be .
RUN chmod +x oneblock_manage_fund_be

# Start a new stage from scratch
FROM golang:1.23.0

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/oneblock_manage_fund_be /app/oneblock_manage_fund_be
COPY config.yaml /app/config.yaml
# Expose port 8080 to the outside world
EXPOSE 8080
# Command to run the executable
CMD ["/app/oneblock_manage_fund_be"]
