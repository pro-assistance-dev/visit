#!/bin/bash

GITHUB_API_REST="pro-assistance-dev/sprob"
last_tag=$(curl -s GET https://api.github.com/repos/pro-assistance-dev/sprob/tags | jq -r '.[].name' | head -n1)

go get github.com/pro-assistance-dev/sprob@${last_tag}
