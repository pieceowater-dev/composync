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

To get started with COMPOSYNC, simply use Docker Compose:

```yaml
services:
  composync:
    image: yurymid/composync
    environment:
      - REPO_URL='https://github.com/pieceowater-dev/lotof.cloud.resources.dev.git' # Your Docker Compose repo
      - BRANCH='main' # Optional: Specify the branch
      - INTERVAL=30 # Check for updates every 30 seconds
      - SCAN-DIR='/' # Optional: Directory with Docker Compose file in remote repo
      - RECURSIVE=true # Optional: Run for every inner Docker Compose file
      - GIT_USERNAME=<YOUR_GIT_USERNAME>
      - GIT_PAT=<YOUR_GIT_PERSONAL_ACCESS_TOKEN>
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - /path/to/your/docker-config.json:/root/.docker/config.json # Mount Docker auth config for pulling images
```

Then, start the service:

```shell
docker-compose up -d
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
   ```shell
   docker run -d --name composync \
       -e INTERVAL=30 \
       -e REPO_URL="https://github.com/pieceowater-dev/lotof.cloud.resources.dev.git" \
       -e GIT_USERNAME=GITHUBUSERNAME \
       -e GIT_PAT=TOKEN \
       -v /var/run/docker.sock:/var/run/docker.sock \
       -v ~/PATH/TO/config.json:/root/.docker/config.json \
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
      - INTERVAL=30
      - REPO_URL='https://github.com/pieceowater-dev/lotof.cloud.resources.dev.git'
      - BRANCH='main'
      - SCAN-DIR='/'
      - RECURSIVE=true
      - GIT_USERNAME=pieceowater
      - GIT_PAT=ghp_12345...abc
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - /path/to/your/docker-config.json:/root/.docker/config.json
```

To run this, navigate to the directory containing your Docker Compose file and execute:

```shell
docker-compose up -d
```

## Running Without Docker

To use COMPOSYNC without Docker:

1. Clone the repository and navigate to the project directory.
2. Run the `index.sh` script with the appropriate parameters:
   ```shell
   ./index.sh \
       --repo="https://github.com/your-repo-url.git" \
       --branch="main" \
       --scan-dir="/" \
       --recursive=true
   ```

COMPOSYNC will then continuously monitor and update your Docker Compose containers.

## License

COMPOSYNC is licensed under the [MIT License](https://github.com/pieceowater-dev/composync/?tab=MIT-1-ov-file).
