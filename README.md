# [COMPOSYNC](https://github.com/pieceowater-dev/composync)
by [pieceowater](https://github.com/pieceowater)

**COMPOSYNC** is a utility designed to keep your Docker Compose containers up to date by continuously syncing them with the latest configurations and dependencies from your remote repositories. It functions similarly to ArgoCD, but for Docker Compose, helping you maintain a single source of truth for your container configurations.

## Key Features

- **Automated Updates**: Automatically detects changes in your remote repository and applies the latest updates.
- **Customizable**: Easily configure the repository URL, branch, scan directory, and other options to fit your needs.
- **Seamless Integration**: Works smoothly with your existing Docker Compose setup.

## How It Works

COMPOSYNC scans a specified directory for Docker Compose files (`docker-compose.yaml` or `docker-compose.yml`). It then pulls the latest changes from the remote repository and applies them to your local environment, ensuring your containers are always synchronized with the latest configurations.

## Quick Start

To get started with COMPOSYNC, simply use Docker:

```bash
docker run -d --name composync \
    -e INTERVAL=10 \ # Interval in minutes to check for updates
    -e REPO_URL="https://github.com/pieceowater-dev/lotof.cloud.resources.dev.git" \ # URL of your remote repository containing Docker Compose files
    -e BRANCH="main" \ # Branch of the repository to use (default is 'main')
    -e SCAN_DIR="/" \ # Directory in the repository where Docker Compose files are located
    -e RECURSIVE=true \ # Whether to search through subdirectories for Docker Compose files
    -e GIT_USERNAME="your_git_username" \ # Your GitHub username for authentication
    -e GIT_PAT="your_personal_access_token" \ # Your GitHub personal access token for authentication
    --volume /var/run/docker.sock:/var/run/docker.sock \ # Mount Docker socket to allow COMPOSYNC to manage Docker containers
    yurymid/composync
```

## Installation Guide

### Prerequisites

Ensure you have the following installed:

- **Docker** (v20.10.x or later)
    - [Docker Installation Guides](https://docs.docker.com/get-docker/)
- **Docker Compose** (v2.14.0 or later) [Optional]
    - [Docker Compose Installation](https://docs.docker.com/compose/install/)

### Configuration Steps

1. **Generate a GitHub Authentication Token**:
   This token is used to authenticate with GitHub and access private repositories if needed.
   ```shell
   echo -n 'GITHUBUSERNAME:TOKEN' | base64
   ```

2. **Create a Configuration File**:
   Create a `config.json` file with the following content, replacing `YOUR_BASE64_ENCODED_TOKEN` with your generated token:
   ```json
   {
     "auths": {
       "ghcr.io": { 
         "auth": "YOUR_BASE64_ENCODED_TOKEN"
       }
     }
   }
   ```

3. **Run the COMPOSYNC Container**:
   Use the following command to run COMPOSYNC with the specified environment variables and volume mounts:
   ```shell
   docker run -d --name composync \
       -e INTERVAL=30 \ # Interval in minutes to check for updates
       -e REPO_URL="https://github.com/pieceowater-dev/lotof.cloud.resources.dev.git" \ # URL of your remote repository
       -e GIT_USERNAME=GITHUBUSERNAME \ # Your GitHub username
       -e GIT_PAT=TOKEN \ # Your GitHub personal access token
       -v /var/run/docker.sock:/var/run/docker.sock \ # Mount Docker socket
       -v ~/PATH/TO/config.json:/root/.docker/config.json \ # Mount Docker auth config
       yurymid/composync
   ```

   > **Note**: For ARM architecture, you may need to add `--platform linux/amd64`.

## Docker Compose Alternative

You can also run COMPOSYNC using Docker Compose. Here’s a sample configuration:

```yaml
services:
  composync:
    image: yurymid/composync
    environment:
      - INTERVAL=30 # Interval in minutes to check for updates
      - REPO_URL=https://github.com/pieceowater-dev/lotof.cloud.resources.dev.git # URL of your remote repository
      - BRANCH=main # Branch of the repository to use
      - SCAN_DIR=/ # Directory in the repository with Docker Compose files
      - RECURSIVE=true # Whether to search through subdirectories
      - GIT_USERNAME=pieceowater # Your GitHub username
      - GIT_PAT=ghp_12345...abc # Your GitHub personal access token
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock # Mount Docker socket
      - /path/to/your/docker-config.json:/root/.docker/config.json # Mount Docker auth config
```

To run this, navigate to the directory containing your Docker Compose file and execute:

```shell
docker-compose up -d
```

## Running Without Docker

To use COMPOSYNC without Docker:

1. **Build and Install COMPOSYNC**:
   Compile and install the COMPOSYNC application:
   ```shell
   go build
   go install
   ```

2. **Run COMPOSYNC**:
   Use the following command to start COMPOSYNC:
   ```shell
   composync go \
       --interval=5 # Interval in minutes to check for updates, or 0 to run once
       --repo=https://github.com/pieceowater-dev/lotof.cloud.resources.dev.git \ # URL of your remote repository
       --branch=main \ # Branch of the repository to use
       --scan-dir=/ \ # Directory with Docker Compose files
       --recursive=true \ # Whether to search through subdirectories
       --username=gitusername \ # Your GitHub username
       --token=gitpat123 # Your GitHub personal access token
   ```

COMPOSYNC will then continuously monitor and update your Docker Compose containers.

## License

COMPOSYNC is licensed under the [MIT License](https://github.com/pieceowater-dev/composync/?tab=MIT-1-ov-file).
