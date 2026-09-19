# imgo

**imgo** is a simple, command-line image processing tool written in Go. It allows you to convert images between different formats (e.g., PNG, JPG, WebP) and compress images efficiently, all from the command line.

## Table of Contents
- [About](#about)
- [Features](#features)
- [Get Started](#get-started)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## About

**imgo** is designed for developers and users who need to quickly convert or compress images directly from the command line. It’s simple, fast, and can be used in scripts or automation tools.

Currently, it supports:
- Converting images between common formats (PNG, JPG, WebP, etc.)
- Compressing images to reduce their file size.

## Features

- [x] Convert images between multiple formats (PNG, JPG, WebP, etc.)
- [x] Compress images with customizable quality
- [x] Batch processing by directory + image type
- [ ] Support for more image formats (future plans)

## Get Started

### Prerequisites

Before you get started, make sure you have the following installed:
- **Go** (v1.16 or later): You can install Go from the [official site](https://golang.org/dl/).
- **Make**: If you don't have `make` installed, you can install it by following the instructions for your system. (It’s available by default on most UNIX-like systems).

### Running the Project

#### 1. Clone the repository

To get started, clone the repository:

```bash
git clone https://github.com/wesleybertipaglia/imgo.git
```

#### 2. Install dependencies and build the project

Navigate to the project folder:

```bash
cd imgo
```

To build the binary, you can use the Makefile. This will handle the building process for you:

```bash
make build
```

> This will create a binary file named imgo.

#### 3. Install globally (optional)
To run imgo from anywhere, you can install it globally:

```bash
make install
```

> This will move the imgo binary to /usr/local/bin or another location in your system’s PATH.

## Usage
Once the binary is installed, you can use the imgo commands from any folder.

`-i` accepts a **single file** or a **directory**. For directories,
also pass the image type (`--from` / `--to`).

Single image (como antes):

```bash
imgo convert -i input.jpg --to png -o ./output
imgo compress -i input.jpg -r 80
```

Whole folder (bulk):

```bash
imgo convert -i ./images --from png --to webp
imgo convert -i ./images --from jpg --to webp -o ./output -r 80
# alias style: --ti instead of --from, -t instead of --to
imgo convert -i ./images --ti png --to webp -o ./output

imgo compress -i ./images --from jpg -r 80
imgo compress -i ./images --from png -r 80 -o ./output
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
