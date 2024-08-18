#!/bin/bash

export RED='\033[31m'
export GREEN='\033[32m'
export YELLOW='\033[33m'
export BLUE='\033[34m'
export NC='\033[0m' # No Color
export BOLD='\033[1m'
export UNDERLINE='\033[4m'
export BLINK='\033[5m'
export INVERT='\033[7m'
export HIDDEN='\033[8m'
export STRIKE='\033[9m'



term_width=$(tput cols)

echo -e "${YELLOW}$(printf "=%.0s" $(seq 1 $term_width))${NC}"
echo -e "${GREEN}${BOLD}--- COMPOSYNC! ---${NC}"
echo -e "${BLUE}COMPOSYNC automates the updating of your Docker Compose containers by continuously pulling and applying the latest changes from your remote repositories' Docker Compose files.${NC}"
echo -e "${BLUE}Enjoy seamless Docker Compose updates!${NC}"
echo -e "by ${BOLD}pieceowater${NC}"
echo -e "${YELLOW}$(printf "=%.0s" $(seq 1 $term_width))${NC}"


# ./index.sh \
#     --repo="https://github.com/pieceowater-dev/lotof.cloud.resources.dev.git" \
#     --branch="main" \
#     --scan-dir="/" \
#     --recursive=true \
#     --username="pieceowater" \
#     --token="ghp_1234567890"

export REPO_URL=""
export BRANCH="main"
export SCAN_DIR="/"
export RECURSIVE=false
export GIT_USERNAME=""
export GIT_PAT=""

# Parse arguments
for arg in "$@"
do
    case $arg in
        --repo=*)
        REPO_URL="${arg#*=}"
        shift
        ;;
        
        --branch=*)
        BRANCH="${arg#*=}"
        shift
        ;;

        --scan-dir=*)
        SCAN_DIR="${arg#*=}"
        shift
        ;;

        --recursive=*)
        RECURSIVE="${arg#*=}"
        shift
        ;;

        --username=*)
        GIT_USERNAME="${arg#*=}"
        shift
        ;;

        --token=*)
        GIT_PAT="${arg#*=}"
        shift
        ;;
        
        *)
        echo -e "${RED}Unknown option $arg${NC}"
        exit 1
        ;;
    esac
done

chmod +x ./store_git_creds.sh
./store_git_creds.sh

echo -e "${GREEN}Repository URL: $REPO_URL${NC}"
echo -e "${GREEN}Branch: $BRANCH${NC}"

if [ -z "$REPO_URL" ]; then
    echo -e "${RED}Error: The --repo parameter is required.${NC}"
    exit 1
fi

mkdir -p PCWT || true
export WORK_DIR="$(pwd)/PCWT"

chmod +x ./setup_workdir.sh

SETUP_OUTPUT=$(./setup_workdir.sh)
SETUP_STATUS=$?

echo -e "Setup output: ${NC}${SETUP_OUTPUT}"

if [ $SETUP_STATUS -ne 0 ]; then
    echo -e "Checking for updates..."
    chmod +x ./fetch_updates.sh
    ./fetch_updates.sh
    UPDATE_STATUS=$?
    if [ $UPDATE_STATUS -eq 1 ]; then
        echo -e "${GREEN}Updates were applied to the repository.${NC}"
    else
        echo -e "${GREEN}No updates were necessary.${NC}"
    fi
else
    echo -e "${GREEN}Repository was just cloned. Skipping update check.${NC}"
fi

chmod +x ./scan_and_apply.sh
./scan_and_apply.sh