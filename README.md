# PR Description AI Generator

CLI tool that uses OpenAI ChatGPT to generate Pull Request text description from git changes. Compares the current branch with a target branch (default: `develop`).

## Requirements

- go
- Open AI API Key

## Install

```shell
go install github.com/carlosflorencio/pr-description-ai
```

## Set the API Key

```shell
# Required shell environment variable
export OPENAI_API_KEY=3RBUSS03KNRSOGPRT847WKT72HHIUKDO
```

## Usage

```shell
Generate PR Description from git changes

Usage:
  pr-description-ai [flags]

Flags:
  -b, --branch string   Target branch to compare against (default "develop")
  -h, --help            help for pr-description-ai
  -m, --model string    OpenAI model to use (default "gpt-4")
```
