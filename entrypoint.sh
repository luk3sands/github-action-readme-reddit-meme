#!/bin/sh -l
GIT_NAME="GitHub Actions"
GIT_EMAIL="github-actions[bot]@users.noreply.github.com"
check_commit() {
local COMMIT_EMAIL="$(git log -1 --pretty=format:'%ae')"
echo $COMMIT_EMAIL
if [ $COMMIT_EMAIL == "$GIT_EMAIL" ]
then
echo "This is a bot commit"
exit 1
fi
}
commit() {
git config --global user.name "$GIT_NAME"
git config --global user.email "$GIT_EMAIL"
git add .
git commit -m "Update README.md"
git push origin master
}
check_commit
commit
git status -s
