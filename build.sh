#!/bin/bash

GIT_COMMIT=$(git rev-parse HEAD)
GIT_COMMIT_SHORT=$(echo $GIT_COMMIT | cut -c 1-7)

# Function to build Docker image
build_docker() {
    service=$1
    echo "Building Docker image for $service..."
    # Change to the service directory and run make
    pushd $service
    make build GIT_COMMIT_SHORT=$GIT_COMMIT_SHORT
    popd
}

# Get a list of changed files from the last commit
CHANGED_DIRS=$(git diff --name-only HEAD^ HEAD | xargs -I {} dirname {} | sort -u)

# Loop through changed directories and build if there's a Makefile inside
for dir in $CHANGED_DIRS; do
    if [ -f "${dir}/Makefile" ]; then
        build_docker $dir
    fi
done

