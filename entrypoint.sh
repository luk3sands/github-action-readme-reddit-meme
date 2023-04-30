#!/bin/sh -l
COMMIT_EMAIL="$(git log -1 --pretty=format:'%ae')"
COMMIT_NAME="$(git log -1 --pretty=format:'%ab')"

git config --global user.email "$GIT_EMAIL"
git config --global user.name "$GIT_NAME"

commit() {
  git add README.md
  git commit -m "Update README.md"
  git push origin master
}

echo "\Updated" >> README.md

commit

git status -s
