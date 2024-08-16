
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

### Running COMPOSYNC

Once Docker and Docker Compose are installed, you can start using COMPOSYNC by running the following command:

```shell
docker run -d --name cmps \
    -e INTERVAL=5 \
    -e REPO_URL="https://github.com/your-repo-url.git" \
    -v /var/run/docker.sock:/var/run/docker.sock composync
```

Make sure to replace the `REPO_URL` with your repository's URL.

## Docker Compose Run

To create a Docker Compose file based on the `docker run` command, you can use the following template:

```yaml
services:
    cmps:
        image: composync
        environment:
            - INTERVAL=30
            - REPO_URL='<your-repo-url>'
            - BRANCH='main'
            - SCAN-DIR='/' # root dir with docker-compose file
            - RECURSIVE=true # if you want to run every inner docker-compose files
        volumes:
            - /var/run/docker.sock:/var/run/docker.sock
```

Replace `<your-repo-url>` with the URL of your repository.

To run the container using the Docker Compose file, navigate to the directory where the file is located and run the following command:

```shell
docker-compose up -d
```

This will start the COMPOSYNC container with the specified environment variables and volume mount.

Remember to adjust any other parameters or configurations according to your specific needs.



---

This addition ensures that users are aware of the minimum requirements before using COMPOSYNC.

## Getting Started (Unconteinerized)

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
