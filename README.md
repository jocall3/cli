# 1231 CLI

The official CLI for the 1231 REST API.

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

```sh
go install 'github.com/stainless-sdks/1231-cli/cmd/1231@latest'
```

### Running Locally

```sh
go run cmd/1231/main.go
```

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
