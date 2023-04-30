#!/bin/sh -l

commit() {
  echo "\Updated" >> README.md
  git add README.md
  git commit -m "Update README.md"
  git push origin master
}

commit

git status -s
