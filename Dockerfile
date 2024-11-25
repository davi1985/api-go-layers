FROM golang:1.23

# set working directory
WORKDIR /go/src/app

# Copy the source code
COPY . .

# Expose the port
EXPOSE 3001

# Build the Go app
RUN go build -o main cmd/main.go

# Run the executable
CMD ["./main"]
