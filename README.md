# README.md Reddit Meme

![example workflow](https://github.com/d3code/github-action-commit-workflow-changes/actions/workflows/action.yaml/badge.svg)

This action gets the latest post with an image from the `r/ProgrammerHumor` subreddit and updates your `README.md`

Place an image with the description `![Reddit]` in your readme and the action will replace it with the latest from the Reddit API and commit it to the branch the action was run on. Make sure you have `Read and write workflow permissions` for the repository you are running this action on.

You can optionally change the commit author and email if you wish to. This will default to:
```yaml
  with:
    commit_author: GitHub Actions
    commit_email: github-actions[bot]@users.noreply.github.com
```

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
        uses: d3code/github-action-commit-workflow-changes@latest
        with:
          commit_author: Luke Sands
          commit_email: luke.sands@d3code.com.au

```

### Example readme

```md
  # Luke Sands
  
  ![Reddit](https://i.redd.it/5izq2xc8j2xa1.jpg)
```

### Workflow permissions

<img src="https://github.com/d3code/github-action-readme-reddit-meme/raw/master/images/menu.png" width="250">

<img src="https://github.com/d3code/github-action-readme-reddit-meme/raw/master/images/setting.png" width="600">

## Example

This image will be updated from r/ProgrammerHumor

![Reddit](https://i.redd.it/5izq2xc8j2xa1.jpg)
