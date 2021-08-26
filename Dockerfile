# Two stage build to create one 
# 1. Start from the golang base image as the builder
FROM golang AS builder

# Set the current workdir inside the cointainer
WORKDIR /go/src/challenge

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached
RUN go mod download

# Mkdir src
RUN mkdir ./src

# Copy the source from the current directory to the Working Directory inside the container
COPY . ./src

# Build the Go app
RUN cd ./src && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o challenge ./cmd

# Copy the built binary to top level
RUN cp ./src/challenge .

# Copy migration file to top level
RUN cp ./src/migrations/init.sql .

# Remove source codes that no longer needed
RUN rm -rf go.* *.go src

# Install sqlite3
RUN apt update && apt install -y apt-transport-https ca-certificates sqlite3
RUN sqlite3 chat.db '.read ./init.sql'

# 2. Use scratch image
FROM scratch

# Set working directory
WORKDIR /root/

# Copy file from builder image
COPY --from=builder /go/src/challenge/challenge .

# Copy database file from builder image
COPY --from=builder /go/src/challenge/chat.db .

# Expose ports
EXPOSE 8080

# Run
CMD ["./challenge"]