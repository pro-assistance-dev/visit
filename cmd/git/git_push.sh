#!/bin/bash

sh ./cmd/git/git_commit.sh $1
if [ $? -gt 0 ]; then
  exit 1
fi

git push -u origin HEAD
