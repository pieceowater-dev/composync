#!/bin/bash

if [ ! -d "$WORK_DIR/.git" ]; then
    echo -e "${GREEN}Setting up the working directory...${NC}"
    cd "$WORK_DIR" || exit
    echo -e "${GREEN}Cloning the repository from $REPO_URL...${NC}"
    git clone "$REPO_URL" .
    echo -e "${GREEN}Repository cloned successfully.${NC}"
    cd ..
    exit 0  # Success
else
    echo -e "${YELLOW}Repository already exists.${NC}"
    exit 1  # Signal that no clone occurred (non-zero for differentiation)
fi