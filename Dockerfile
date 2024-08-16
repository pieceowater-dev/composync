FROM alpine:latest

# Install dependencies and Docker Compose in a single RUN command
RUN apk add --no-cache \
    git \
    curl \
    docker \
    bash

# Install Docker Compose
RUN curl -L "https://github.com/docker/compose/releases/download/v2.14.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && \
    chmod +x /usr/local/bin/docker-compose

# Set the working directory
WORKDIR /app

# Copy application files
COPY . /app/

# Set environment variables
ENV INTERVAL=5 \
    REPO_URL="" \
    BRANCH="main" \
    SCAN_DIR="/" \
    RECURSIVE=false

# Copy and make the entrypoint script executable
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Mount the Docker socket volume
VOLUME /var/run/docker.sock

# Use the entrypoint script to start the process
ENTRYPOINT ["/entrypoint.sh"]