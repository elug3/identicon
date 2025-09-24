# identicon
identicon generator written in Go

A simple command-line tool that generates unique identicon images from input strings. Identicons are visual representations of a hash value, usually used as default profile pictures.

## Features

- Generates 5x5 grid identicon patterns
- Symmetric design for visual appeal
- Unique colors based on input hash
- PNG output format (250x250 pixels)
- Command-line interface for easy usage

## Installation

```bash
go build -o identicon .
```

## Usage

```bash
# Basic usage with default output filename
./identicon -input=john.doe@example.com

# Custom output filename
./identicon -input=myusername -output=avatar.png

# Show help
./identicon -help
```

## Examples

Generate an identicon for an email address:
```bash
./identicon -input=user@example.com
# Creates: user@example.com.png
```

Generate an identicon with custom filename:
```bash
./identicon -input=myusername -output=my-avatar.png
# Creates: my-avatar.png
```

## How it works

1. Takes an input string and computes its MD5 hash
2. Uses the first 3 bytes of the hash to determine RGB color values
3. Uses subsequent bytes to create a symmetric 5x5 grid pattern
4. Renders the pattern as a 250x250 pixel PNG image

Each unique input string will always generate the same identicon, making it perfect for consistent user avatars.
