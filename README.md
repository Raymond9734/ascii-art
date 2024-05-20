# ASCII Art Banner Generator

This Go program generates ASCII art banners from input text, optionally applying color to specified characters. It supports multiple font styles and allows for custom colorization.

## Features

- Generate ASCII art using different fonts.
- Apply colors to specified characters in the output.
- Handle special characters and newlines.

## Requirements

- Go 1.16 or higher

## Installation

1. Clone the repository:
    ```sh
    git clone https://learn.zone01kisumu.ke/git/rcaleb/ascii-art.git
    ```
2. Navigate to the project directory:
    ```sh
    cd ascii-art/Color
    ```

## Usage

```sh
go run . [OPTION] [STRING]

go run . --color=red "He" "Hello World"

```

### Examples


1. Generate an ASCII art banner with color:
    ```sh
    go run . --color=red H "Hello, World!"
    ```

### Options

- `--color=<color>`: Specifies the color to be applied to the characters in the STRING. Replace `<color>` with the desired color name.
- `standard`, `thinkertoy`, `shadow`: Specifies the font style to be used for the ASCII art.

## How It Works

The program processes command-line arguments to determine the font style, color, and input string. It reads the appropriate font file and generates the ASCII art line by line. If a color option is provided, it applies the specified color to the designated characters.

### Code Overview

- `main()`: Entry point of the program. Parses arguments, processes the input string, and generates the ASCII art.
- `GetLine(num int, filename string) string`: Reads the specified line from the font file.
- `ToColor(str1 string, str2 string) bool`: Determines if the characters in `str2` should be colored in `str1`.

## Directory Structure

```
asci-art/
├── BannerFiles/
│   ├── standard.txt
│   ├── thinkertoy.txt
│   └── shadow.txt
├── Color/
    ├── Functions/
    │   └── color.go 
    └── main.go
```

- `BannerFiles/`: Contains the font files.
- `Color/Functions/`: Contains color-related functions.
- `Banner/`: Contains banner generation functions.
- `main.go`: Main program file.

## License

This project is licensed under the MIT License.

## Author

Your Name - [your-email@example.com](mailto:your-email@example.com)

Feel free to reach out if you have any questions or suggestions!