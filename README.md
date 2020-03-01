# Go GitHub Stats

A Go app that hits the GitHub REST API and creates a report on things that matter to you or your development team.

## Motivations

1. Learn Go with a fun little project!
2. Create an app that will periodically give my team at work insights into all of the different repositories we have ownership over
3. Learn how to build messages for Slack
4. Stay in-shape with AWS

## Requirements

- Golang 1.14

## Setup

Modify `config.json` to setup the repositories you want a report on. The provided config has a duplicate repository definition to make it more obvious that you can configure multiple repositories at once.

## Running

If you're in the top directory of the repo and you're trying to run the application, try a classic `go run`.

## TO-DO/Goals
- [ ] Create Markdown report builder
- [ ] Add Slack messaging
- [ ] Switch from GitHub v3 API to the GraphQL API
