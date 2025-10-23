# 1 Use official lightweight Go image
FROM golang:1.25.1-alpine AS builder

# 2 Set working directory inside container
WORKDIR /app

# 3 Copy go.mod and go.sum first (to leverage caching)
COPY go.mod go.sum ./

# 4 Download dependencies 
RUN go mod download

# 5 Copy the rest of the source code
COPY . . 

# 6 Build the Go binary (optimized)
RUN go build -o notes-api main.go

# 7 Use a smaller image for production
FROM alpine:latest

WORKDIR /root/

# 8 Copy binary from builder stage
COPY --from=builder /app/notes-api .

# 9 Expose port 8080 for HTTP traffic
EXPOSE 8080

#10 Run the application
CMD ["./notes-api"]
