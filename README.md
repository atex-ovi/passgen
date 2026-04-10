<div align="center">

# PassGen - Simple Password Generator

![License](https://img.shields.io/badge/license-MIT-green)
![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-blue)

A simple, fast, and lightweight password generator CLI tool written in Go. Generate secure random passwords with custom length, special characters, and numbers.

</div>

## Features

- Generate random passwords with custom length
- Option to include/exclude special characters
- Option to include/exclude numbers
- Simple and intuitive CLI interface
- Cross-platform (Linux, macOS, Windows)
- No external dependencies

## Installation

### Option 1: Go Install

```bash
go install github.com/atex-ovi/passgen@latest
```

### Option 2: Build from Source

```bash
git clone https://github.com/atex-ovi/passgen.git
cd passgen
go build -o passgen main.go
sudo mv passgen /usr/local/bin/  # Linux/macOS
```

### Option 3: Download Binary

Download the pre-built binary for your platform from [Releases](https://github.com/atex-ovi/passgen/releases).

## Usage

### Basic Usage

```bash
# Generate default password (12 characters)
passgen

# Generate 16 characters password
passgen -l 16

# Generate with special characters
passgen -s

# Generate 20 characters with special chars
passgen -l 20 -s

# Generate without numbers
passgen -n false
```

### Command Line Options

| Flag | Default | Description |
|------|---------|-------------|
| `-l` | 12 | Password length |
| `-s` | false | Include special characters (!@#$%^&*()_+-=[]{}) |
| `-n` | true | Include numbers (0-9) |
| `-h` | false | Show help message |

### Examples

```bash
$ passgen -l 20 -s

🔐 Generated Password:
━━━━━━━━━━━━━━━━━━━━━━
  P#m5&jc*ni[wK9fOk!X
━━━━━━━━━━━━━━━━━━━━━━
  Length: 20 | Special: true | Numbers: true
```

## Development

### Prerequisites

- Go 1.21 or higher

### Run Locally

```bash
git clone https://github.com/atex-ovi/passgen.git
cd passgen
go run main.go -l 16 -s
```

## License

MIT License

Copyright (c) 2026 Atex Ovi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

---

<div align="center">

**Simple, fast, and secure password generator — 100% open source**

</div>
