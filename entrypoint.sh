#!/bin/sh
exec /usr/local/bin/composync go \
    --interval="${INTERVAL}" \
    --repo="${REPO_URL}" \
    --branch="${BRANCH}" \
    --scan-dir="${SCAN_DIR}" \
    --recursive="${RECURSIVE}" \
    --username="${GIT_USERNAME}" \
    --token="${GIT_PAT}"