#!/bin/sh -l

commit() {
  echo "\nUpdated" >> README.md
  git add README.md
  git commit -m "Update README.md"
  git push origin master
}

commit

git status -s
