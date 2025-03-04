#!/bin/bash

# only used for deps.sh

# Function to fetch commit hash from a given Chromium tag
get_commit_from_tag() {
    local tag="$1"
    local url="https://chromium.googlesource.com/chromium/src.git/+/${tag}?format=JSON"

    # Fetch data using curl
    response=$(curl -s "$url")

    # Remove Gitiles security prefix and extract commit hash
    commit_hash=$(echo "$response" | sed '1d' | jq -r '.commit')

    if [[ "$commit_hash" == "null" ]]; then
        ecit 1
    else
        echo "$commit_hash"
    fi
}

# Check if tag is provided
if [[ $# -ne 1 ]]; then
    echo "Usage: $0 <tag>"
    exit 1
fi

get_commit_from_tag "$1"
