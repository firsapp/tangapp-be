# Start with the official Golang image
FROM golang:1.20 as builder

# Set up the working directory and copy project files
WORKDIR /app
COPY . .

# Download dependencies and build the application
RUN go mod download
RUN go build -o server .

# Start with a minimal base image for the final container
FROM gcr.io/distroless/base-debian10

# Copy the binary from the builder stage
COPY --from=builder /app/server /server

# Expose the port
EXPOSE 8080

# Run the application
CMD ["/server"]
