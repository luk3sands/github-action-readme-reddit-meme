# README.md Reddit Meme

![example workflow](https://github.com/d3code/github-action-commit-workflow-changes/actions/workflows/action.yaml/badge.svg)

This action gets the latest post with an image from the `r/ProgrammerHumor` subreddit and updates your `README.md`

Place an image with the description `![Reddit]` in your readme and the action will replace it with the latest from the Reddit API and commit it to the branch the action was run on. Make sure you have `Read and write workflow permissions` for the repository you are running this action on.

### Example workflow

Updates at 6am every day and on push to master

```yaml
name: Reddit Meme

on:
  push:
    branches: [ master ]
  schedule:
    - cron: '0 6 * * *'

jobs:
  deploy:
    name: Update Reddit
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Update Reddit
        uses: d3code/github-action-commit-workflow-changes@master
        with:
          commit_author: GitHub Actions
          commit_email: <github-actions[bot]@users.noreply.github.com>

```

### Example readme

```md
# Luke Sands

![Reddit](https://i.redd.it/gasrh7d2h1xa1.png)
```

### Workflow permissions

<img src="https://github.com/d3code/github-action-readme-reddit-meme/raw/master/images/menu.png" width="250">

<img src="https://github.com/d3code/github-action-readme-reddit-meme/raw/master/images/setting.png" width="600">

## Example

This image will be updated from r/ProgrammerHumor

![Reddit](https://i.redd.it/gasrh7d2h1xa1.png)
