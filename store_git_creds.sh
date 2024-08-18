#!/bin/bash

# Store the credentials using Git credential helper
git config --global credential.helper store

# Correctly format and store credentials using the credential helper
echo -e "protocol=https\nhost=github.com\nusername=$GIT_USERNAME\npassword=$GIT_PAT" | git credential approve