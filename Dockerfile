FROM golang:alpine as builder

# Set the Current Working Directory to /app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Create a new stage for the final image
FROM alpine:latest 

# Set the Current Working Directory to /app
WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Command to run on container start
CMD ["./main"]

# Expose port 8080 to the docker host, so we can access it from the outside.
EXPOSE 8080
