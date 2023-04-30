# README.md Reddit Meme

![example workflow](https://github.com/d3code/github-action-commit-workflow-changes/actions/workflows/action.yaml/badge.svg)

This action gets the latest post with an image from the /r/ProgrammerHumor subreddit and updates your `README.md`

Place an image with the description `![Reddit]` in your readme and the action will replace it with the latest from the Reddit API and commit it to the branch the action was run on. Make sure you have `Read and write workflow permissions` for the repository you are running this action on.

You can find this setting for your repository here:
- Settings
- Actions > General
- Workflow permissions
- Read and write workflow permissions

Add `![Reddit](https://i.redd.it/gasrh7d2h1xa1.png)` to your README.md and set a cron schedule for your action and get a new meme at your desrired schedule.

![Reddit](https://i.redd.it/gasrh7d2h1xa1.png)
