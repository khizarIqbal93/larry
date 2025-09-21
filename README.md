# Larry

A simple Go tool for Git version management and semantic versioning operations.

## Overview

Larry is a command-line utility that helps developers manage Git tags and version numbers. It provides functionality to:

- Retrieve the last commit hash from your Git repository
- Parse and manage semantic version tags (v1.0.0 format)
- Increment version numbers (major, minor, or patch)

## Features

- **Git Integration**: Automatically reads from your local Git repository
- **Semantic Versioning**: Supports standard semantic versioning format (v1.0.0)
- **Version Incrementing**: Increment major, minor, or patch versions
- **Commit Hash Retrieval**: Get the short hash of the last commit

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

```bash
go install github.com/kdawg500/larry@latest
```

## Usage

### Basic Usage

Run Larry from within a Git repository:

```bash
./larry
```

This will output:

- The incremented patch version of the last tag
- The short hash of the last commit

### Example Output

```
v1.0.1
a1b2c3
```

### Running Tests

```bash
go test ./...
```

### Project Structure

```
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
