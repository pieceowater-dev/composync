#!/bin/bash

cd "$WORK_DIR" || exit

REMOTE_COMMIT=$(git ls-remote "$REPO_URL" "$BRANCH" | awk '{print $1}')

LOCAL_COMMIT=$(git rev-parse "$BRANCH")

if [ "$REMOTE_COMMIT" != "$LOCAL_COMMIT" ]; then
    echo -e "${GREEN}Changes detected in the remote repository.${NC}"
    echo -e "${BLUE}Pulling the latest changes...${NC}"
    git pull origin "$BRANCH"
    exit 1  # updates were found and pulled
else
    echo -e "${YELLOW}No changes detected in the remote repository.${NC}"
    exit 0  # no updates were found
fi