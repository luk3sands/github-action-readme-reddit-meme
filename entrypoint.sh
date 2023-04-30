#!/bin/sh -l
GIT_NAME="GitHub Actions"
GIT_EMAIL="github-actions[bot]@users.noreply.github.com"

commit() {
  git add README.md
  git commit -m "Update README.md"
  git push origin master
}

echo "\nUpdated README.md" >> README.md

commit

git status -s
