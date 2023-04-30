#!/bin/sh -l

commit() {
  git add README.md
  git commit -m "Update README.md"
  git push origin master
}

echo "\nUpdated README.md" >> README.md

commit

git status -s
