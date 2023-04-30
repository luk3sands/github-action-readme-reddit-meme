#!/bin/bash

GIT_NAME="GitHub Actions"
GIT_EMAIL="github-actions[bot]@users.noreply.github.com"

check_commit() {
  local COMMIT_EMAIL="$(git log -1 --pretty=format:'%ae')"
  echo $COMMIT_EMAIL
  
  if [ $COMMIT_EMAIL == $GIT_EMAIL ];
  then
      echo "This is a bot commit"
      exit 1
  fi
}
