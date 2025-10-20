# Larry

<!-- markdownlint-disable-next-line MD033 -->
<img src="https://i.redd.it/ty3qlu4swjle1.jpeg" alt="A picture of Larry intently into your soul" width="220" />

A simple Go tool for Git version management and semantic versioning operations.

## Overview

Larry is a command-line utility that helps developers manage Git tags and version numbers. It provides functionality to:

- Retrieve the last commit hash from your Git repository
- Parse and manage semantic version tags (v1.0.0 format)
- Increment version numbers (major, minor, or patch)
- Automatically create annotated Git tags with incremented versions

## Features

- **Git Integration**: Automatically reads from your local Git repository
- **Semantic Versioning**: Supports standard semantic versioning format (v1.0.0)
- **Version Incrementing**: Increment major, minor, or patch versions
- **Automatic Tag Creation**: Creates annotated Git tags with incremented versions
- **Commit Hash Retrieval**: Uses the current commit hash for tag creation
- **Custom Tag Messages**: Add descriptive messages to your version tags

## Installation

### Prerequisites

- Go 1.24.1 or later
- Git repository (for Git operations)

### Build from Source

```bash
git clone https://github.com/kdawg500/larry.git
cd larry
go build -o larry main.go
```

### Install Globally

Make sure your GOBIN and PATH variables are set up correctly:

Add this to your .zshrc/.bashrc

```sh
export PATH="$PATH:$HOME/go/bin"
```

```bash
go install github.com/kdawg500/larry@latest
```

## Usage

### Basic Usage

Run Larry from within a Git repository to increment version numbers and create Git tags:

```bash
larry increment <major|minor|patch>
```

### Options

- `-m <message>`: Add a custom message to the created tag

### Examples

Increment the patch version (e.g., v1.0.0 → v1.0.1):

```bash
larry increment patch
```

Increment the minor version (e.g., v1.0.0 → v1.1.0):

```bash
larry increment minor
```

Increment the major version (e.g., v1.0.0 → v2.0.0):

```bash
larry increment major
```

Create a tag with a custom message:

```bash
larry -m "Bug fixes and improvements" increment patch
```

### What Larry Does

When you run Larry, it will:

1. Get the last Git tag from your repository
2. Parse the semantic version from that tag
3. Increment the specified version component (major, minor, or patch)
4. Get the current commit hash
5. Create a new annotated Git tag with the incremented version
6. Output the new version string

### Example Output

```text
v1.0.1
```

### Running Tests

```bash
go test ./...
```

### Project Structure

```text
larry/
├── main.go              # Main application entry point
├── parse/
│   ├── gittag.go        # Git tag parsing and version management
│   ├── gittag_test.go   # Tests for tag functionality
│   └── gitlog.go        # Git log parsing
├── go.mod               # Go module definition
└── go.sum               # Go module checksums
```

## License

This project is open source and available under the [MIT License](LICENSE).

## Author

Created by [kdawg500](https://github.com/kdawg500)
