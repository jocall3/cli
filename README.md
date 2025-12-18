# Jocall3 CLI

The official CLI for the Jocall3 REST API.

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

<!-- x-release-please-start-version -->

```sh
go install 'github.com/jocall3/1231-cli/cmd/jocall3@latest'
```

### Running Locally

<!-- x-release-please-start-version -->

```sh
go run cmd/jocall3/main.go
```

<!-- x-release-please-end -->

## Usage

The CLI follows a resource-based command structure:

```sh
jocall3 [resource] [command] [flags]
```

```sh
jocall3 users register \
  --email alice.w@example.com \
  --name 'Alice Wonderland' \
  --password 'SecureP@ssw0rd2024!'
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
