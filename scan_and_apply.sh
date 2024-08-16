#!/bin/bash
echo -e "${BLINK}Scanning the directory: ${UNDERLINE}$SCAN_DIR${NC}"
echo -e "Recursive: ${UNDERLINE}$RECURSIVE${NC}"

SCAN_PATH="${WORK_DIR}${SCAN_DIR}"

if [ "$RECURSIVE" = true ]; then
    FIND_OPTIONS=""
else
    FIND_OPTIONS="-maxdepth 1"
fi

find "$SCAN_PATH" $FIND_OPTIONS -type f \( -name 'docker-compose.yaml' -o -name 'docker-compose.yml' \) | while read -r file; do
    echo -e "${GREEN}Found${NC} docker-compose file: ${BLINK}${UNDERLINE}$file${NC}"
    docker-compose -f "$file" up -d || docker compose -f "$file" up -d || echo -e "${RED}Error running docker-compose up -d${NC}"
done