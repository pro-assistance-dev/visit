#!/bin/bash

sh ./cmd/prepare_git_hooks.sh

commit_message=$1

if [ -z "$commit_message" ]; then
  echo "No commit message supplied! Commit stopped."
  exit 1
fi

git pull origin develop
git add .
git commit -m "$commit_message"
