#!/bin/bash

# Ensure that the script is executable
chmod +x /app/index.sh

# Log that the script is starting
echo "$(date) Starting composync with an interval of ${INTERVAL} seconds."

# Run the task in an infinite loop with the specified interval
while true; do
    echo "$(date) Running composync at $(date)"
    
    # Run the index.sh script and output its logs to both console and log file
    /app/index.sh --repo=${REPO_URL} --branch=${BRANCH} --scan-dir=${SCAN_DIR} --recursive=${RECURSIVE} --username=${GIT_USERNAME} --token=${GIT_PAT} 2>&1 | tee -a /var/log/composync.log
    
    # Wait for the specified interval before the next run
    sleep ${INTERVAL}
done