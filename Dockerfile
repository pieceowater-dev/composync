# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache \
    git \
    curl \
    bash \
    ncurses

# Copy Go module files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o composync ./main.go

# Stage 2: Create the final image with runtime dependencies
FROM alpine:latest

# Create a non-root user and group
RUN addgroup -S composync && adduser -S composync -G composync

# Install runtime dependencies including bash, git, and Docker CLI
RUN apk add --no-cache \
    bash \
    curl \
    git \
    ncurses \
    docker-cli

# Install Docker Compose
RUN curl -L "https://github.com/docker/compose/releases/download/v2.14.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && \
    chmod +x /usr/local/bin/docker-compose

# Change permissions of Docker socket to allow access to the docker group
RUN chown root:docker /var/run/docker.sock && chmod 660 /var/run/docker.sock

# Copy the built Go application from the builder stage
COPY --from=builder /app/composync /usr/local/bin/composync

# Change ownership of the application to the composync user
RUN chown composync:composync /usr/local/bin/composync

# Copy the entrypoint script, set permissions, and change ownership
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh && \
    chown composync:composync /usr/local/bin/entrypoint.sh

# Add composync to the docker group for access to the socket
RUN addgroup composync docker

# Switch to the composync user
USER composync

# Set the entrypoint for the container
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

# Default command (uses environment variables)
CMD ["go", "--interval=${INTERVAL}", "--repo=${REPO_URL}", "--branch=${BRANCH}", "--scan-dir=${SCAN_DIR}", "--recursive=${RECURSIVE}", "--username=${GIT_USERNAME}", "--token=${GIT_PAT}"]