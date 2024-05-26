# GitHub Actions

Continuous integration and continuous delivery  (CI/CD) platform that allows you to automate your build, test, and deployment pipeline.

## Components

- Workflows: Configurable automated process that will run one or more jobs.

- Events: Specific activity in a repository that triggers a workflow.

- Jobs: Set of steps in a workflow that is executed on the same runner.

- Actions: Custom application for the GitHub Actions platform that performs a complex but frequently repeated task.

- Runners: Server that runs your workflows when they are triggered.

## Example workflow

GitHub Actions use YAML syntax to define the workflow. Each workflow is stored as a separate YAML file in your code repository, in a directory named `.github/workflows`.

```yaml
name: learn-github-actions
run-name: ${{ github.actor }} is learning GitHub Actions
on: [push]
jobs:
  check-bats-version:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '20'
      - run: npm install -g bats
      - run: bats -v    
```

In this workflow, GitHub Actions checks out the pushed code, installs the bats testing framework, and runs a basic command to output the bats version: `bats -v`.
