# imgool

**imgool** is a simple, command-line image processing tool written in Go. It allows you to convert images between different formats (e.g., PNG, JPG, WebP) and compress images efficiently, all from the command line.

## Table of Contents
- [About](#about)
- [Features](#features)
- [Get Started](#get-started)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## About

**imgool** is designed for developers and users who need to quickly convert or compress images directly from the command line. It’s simple, fast, and can be used in scripts or automation tools.

Currently, it supports:
- Converting images between common formats (PNG, JPG, WebP, etc.)
- Compressing images to reduce their file size.

## Features

- [x] Convert images between multiple formats (PNG, JPG, WebP, etc.)
- [x] Compress images with customizable quality
- [ ] Support for more image formats (future plans)
- [ ] Batch processing for multiple images (future plans)

## Get Started

### Prerequisites

Before you get started, make sure you have the following installed:
- **Go** (v1.16 or later): You can install Go from the [official site](https://golang.org/dl/).
- **Make**: If you don't have `make` installed, you can install it by following the instructions for your system. (It’s available by default on most UNIX-like systems).

### Running the Project

#### 1. Clone the repository

To get started, clone the repository:

```bash
git clone https://github.com/wesleybertipaglia/imgool.git
```

#### 2. Install dependencies and build the project

Navigate to the project folder:

```bash
cd imgool
```

To build the binary, you can use the Makefile. This will handle the building process for you:

```bash
make build
```

> This will create a binary file named imgool.

#### 3. Install globally (optional)
To run imgool from anywhere, you can install it globally:

```bash
make install
```

> This will move the imgool binary to /usr/local/bin or another location in your system’s PATH.

## Usage
Once the binary is installed, you can use the imgool commands from any folder.

For example, to convert an image:

```bash
imgool convert -i input.jpg -t png -o ./output
```

To compress an image:

```bash
imgool compress -i input.jpg -r 80
```

### Using the Makefile

```bash
make dev              - Runs the application with 'go run main.go'
make build            - Builds the binary executable
make clean            - Removes the built binary
make install          - Installs the binary globally
make start            - Runs the build
make fmt              - Formats the Go code
make lint             - Lint the Go code
```

## Contributing

Contributions are welcome! If you have any suggestions or improvements, please open an issue or a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
