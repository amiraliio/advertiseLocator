# Start from the latest golang base image
FROM golang:1.13.4 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o application

######## Start a new stage from scratch #######
FROM alpine:3.10.3

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy config file to the build project directory
COPY --from=builder /app/config.yaml .

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/application .

# Expose port 3749 to the outside world
EXPOSE 3749

# Build Args
ARG STORAGE_DIR=/root/storage

# Create Log Directory
RUN mkdir -p ${STORAGE_DIR}

# Declare volumes to mount
VOLUME [${STORAGE_DIR}]

# Command to run the executable
CMD ["./application"]