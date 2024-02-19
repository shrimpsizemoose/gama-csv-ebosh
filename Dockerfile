# Use a multi-stage build to keep the final image small
# Build stage
FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.source https://github.com/shrimpsizemoose/gama-csv-ebosh-5678

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 5678 to the outside world
EXPOSE 5678

# Command to run the executable
CMD ["./main"]

