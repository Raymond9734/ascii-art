

# ASCII Art Printer

This Go program reads a text input and prints its ASCII art representation using a predefined set of ASCII characters stored in a file named `standard.txt`.

## Features

- Converts text input to ASCII art
- Supports newline characters (`\n`)
- Handles missing characters with an error message

## Prerequisites

- Go installed on your machine

## Installation

1. Clone the repository:

    ```bash
    git clone https://learn.zone01kisumu.ke/git/rcaleb/ascii-art.git
    ```

2. Navigate to the project directory:

    ```bash
    cd ascii-art
    ```



## Usage

Run the program with a text input as a command-line argument:

```bash
go run .  "Hello, World!" | cat -e
```



## Example

```bash
go run .  "Hello, World!" | cat -e

```  
## How It Works


1. **Input Parsing**: 
    - The program takes a text input from the command-line arguments using `os.Args[1]`.
    - The input may contain newline characters (`\n`), which are replaced with actual newline characters using `strings.Replace`.

2. **ASCII Art Mapping**:
    - The `Banner` function reads a file named `standard.txt` to load predefined ASCII art characters.
    - It splits the file content into lines and maps each character to its corresponding ASCII art representation.

3. **Text to ASCII Conversion**:
    - The `PrintAscii` function processes each character in the input text.
    - For each character, it looks up its ASCII art representation from the map created by the `Banner` function.
    - The ASCII art lines for each character are collected in a 2D slice (`Output`).
  
4. **Printing ASCII Art**:
    - Finally, the `Output` slice is printed to the console, displaying the input text as ASCII art.
  

### Example:

For the input text `"Hello\nWorld"`, the program will:

1. Replace `\n` with actual newline characters.
2. Split the text into lines: `["Hello", "World"]`.
3. Look up each character's ASCII art representation from `standard.txt`.
4. Print the ASCII art representation of each character to the console, maintaining the line structure.

```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
__          __                 _       _  
\ \        / /                | |     | | 
 \ \  /\  / /    ___    _ __  | |   __| | 
  \ \/  \/ /    / _ \  | '__| | |  / _` | 
   \  /\  /    | (_) | | |    | | | (_| | 
    \/  \/      \___/  |_|    |_|  \__,_| 
                                          

                                                                        
```


### Custom Banner Characters

You can provide a custom banner file to use for printing. The file should contain characters represented as ASCII art with each character separated by newline characters. The program reads 8 lines for each character.

To use a custom banner file, execute the following command:

```bash

go run . "Your text" "customBannerFileName"

```

`Replace customBannerFileName with the name of your custom banner file without the ./ and  .txt .`


## File Structure

- `main.go`: Main source code file
- `standard.txt`: File containing ASCII art characters
- `shadow.txt` : File containing ASCII art characters
- `thinkertoy.txt` : File containing ASCII art characters
- `Banner folder `: contains the LoadBanner, PrintBanner, RemoveSpecialcase and filecheck  functions 

## Contributing

Feel free to contribute to this project. Create a pull request with your changes, or open an issue if you find any bugs or have suggestions.

