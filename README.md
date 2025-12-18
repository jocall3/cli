# 1231 CLI

The official CLI for the 1231 REST API.

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

<!-- x-release-please-start-version -->

```sh
go install 'github.com/jocall3/cli/cmd/1231@latest'
```

### Running Locally

<!-- x-release-please-start-version -->

```sh
go run cmd/1231/main.go
```

<!-- x-release-please-end -->

## Usage

The CLI follows a resource-based command structure:

```sh
1231 [resource] [command] [flags]
```

```sh
1231 users register \
  --email alice.w@example.com \
  --name 'Alice Wonderland' \
  --password 'SecureP@ssw0rd2024!'
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
