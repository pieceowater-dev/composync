# COMPOSYNC

COMPOSYNC is a powerful utility that automates the updating of your Docker Compose containers. It continuously pulls and applies the latest changes from your remote repositories' Docker Compose files, ensuring that your containers are always up to date.

## How it Works

COMPOSYNC scans a specified directory for Docker Compose files (`docker-compose.yaml` or `docker-compose.yml`). It then pulls the latest changes from the remote repository and applies them to your local environment. This process ensures that your containers are synchronized with the latest configurations and dependencies.

## Features

- **Automated Updates**: COMPOSYNC automatically detects changes in the remote repository and pulls the latest updates.
- **Flexible Configuration**: You can specify the repository URL, branch, scan directory, and recursion options to customize the synchronization process.
- **Easy Integration**: COMPOSYNC seamlessly integrates with your existing Docker Compose workflow, allowing you to effortlessly keep your containers up to date.

## Installation Guide

Before you begin using COMPOSYNC, ensure that you have the following installed:

### Docker
- **Version:** 20.10.x or later

You can install Docker by following the official installation guide for your operating system:
- [Docker for Linux](https://docs.docker.com/engine/install/)
- [Docker for Windows](https://docs.docker.com/desktop/install/windows-install/)
- [Docker for macOS](https://docs.docker.com/desktop/install/mac-install/)

### Docker Compose (Optional)
- **Version:** 2.14.0 or later

You can install Docker Compose by following the official installation guide:
- [Docker Compose Installation](https://docs.docker.com/compose/install/)

### Configuration

1. **Generate Authentication Token**: Execute the following command to generate a base64-encoded GitHub authentication token:
   ```shell
   echo -n 'GITHUBUSERNAME:TOKEN' | base64
   ```

2. **Create Configuration File**: Create a file named `config.json` with the following content, replacing `YOUR_BASE64_ENCODED_TOKEN` with the output from the previous step:
   ```json
   {
       "auths": {
           "ghcr.io": { 
               "auth": "YOUR_BASE64_ENCODED_TOKEN"
           }
       }
   }
   ```

3. **Run COMPOSYNC Container**: Start the COMPOSYNC container using the following command, replacing placeholders with your actual values:
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

   also you can run with this flag ```--platform linux/amd64```

## Docker Compose Run

To create a Docker Compose file based on the `docker run` command, you can use the following template:

```yaml
services:
    cmps:
        image: composync
        environment:
            - INTERVAL=30
            - REPO_URL='https://github.com/pieceowater-dev/lotof.cloud.resources.dev.git'
            - BRANCH='main'
            - SCAN-DIR='/' # root dir with docker-compose file
            - RECURSIVE=true # if you want to run every inner docker-compose file
            - GIT_USERNAME=pieceowater
            - GIT_PAT=ghp_12345...abc
        volumes:
            - /var/run/docker.sock:/var/run/docker.sock
            - /path/to/your/docker-config.json:/root/.docker/config.json
```

To run the container using the Docker Compose file, navigate to the directory where the file is located and run the following command:

```shell
docker-compose up -d
```

This will start the COMPOSYNC container with the specified environment variables and volume mount.

Remember to adjust any other parameters or configurations according to your specific needs.

## Getting Started (Uncontainerized)

To start using COMPOSYNC, follow these steps:

1. Clone the repository and navigate to the project directory.
2. Run the `index.sh` script with the appropriate parameters to configure COMPOSYNC.
3. Sit back and let COMPOSYNC handle the rest! It will continuously monitor and update your Docker Compose containers.

## Requirements

- Docker
- Git

## Usage

```shell
./index.sh \
    --repo="https://github.com/your-repo-url.git" \
    --branch="main" \
    --scan-dir="/" \
    --recursive=true
```

## License

COMPOSYNC is released under the [MIT License](LICENSE).
