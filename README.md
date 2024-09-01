# Password Generator

## Overview

`password_generator` is a command-line tool designed to generate random passwords with a variety of customizable options. This tool is ideal for creating secure passwords that meet specific requirements, whether for personal use, development, or deployment.

## Features

- **Length Specification**: Set the exact length of the generated password.
- **Inclusion of Symbols**: Optionally include special symbols to enhance password complexity.
- **Inclusion of Numbers**: Optionally include numbers to meet common security policies.
- **Exclusion of Similar Characters**: Optionally exclude characters that can be easily confused (e.g., `1`, `l`, `I`, `0`, `O`).
- **Length Constraints**: Define minimum and maximum length constraints for the password.
- **Verbose Output**: Get detailed information about the generation process.
- **Configuration File Support**: Placeholder for configuration file support (for future enhancement).

## Installation

To use this password generator, you'll need to have Go installed on your machine. You can download Go from the [official Go website](https://golang.org/dl/).

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/JakubSchwenkbeck/password-generator.git
   ```
2. **Navigate to the Project Directory:**

```bash
cd password-generator
```
3. **Build the Project:**

```bash
go build -o password_generator
```
4.** Run the Tool:**

After building the project, you can run the tool using the following command:

```bash
./password_generator [flags]
```
## Usage
You can use the password generator tool from the command line with various flags to customize the generated password. Here are some example commands and their descriptions:

### Generate a password with default settings:

```bash
./password_generator
```
Generate a password of length 16 including symbols:

```bash
./password_generator --length 16 --symbols
```
Generate a password with numbers, excluding similar characters, and verbose output:

```bash
./password_generator --numbers --exclude-similar --verbose
```
Specify minimum and maximum length constraints:

```bash
./password_generator --min-length 12 --max-length 20
```

## Flags

The `password_generator` tool supports the following command-line flags:

- `-l, --length int`  
  **Length of the password**  
  The number of characters in the generated password. Default is 12.

- `-s, --symbols`  
  **Include symbols**  
  Whether to include special symbols in the password.

- `-n, --numbers`  
  **Include numbers**  
  Whether to include numbers in the password.

- `-x, --exclude-similar`  
  **Exclude similar characters**  
  Whether to exclude characters that are easily confused with each other (e.g., `1`, `l`, `I`, `0`, `O`).

- `--min-length int`  
  **Minimum length of the password**  
  The minimum number of characters for the password. Default is 8.

- `--max-length int`  
  **Maximum length of the password**  
  The maximum number of characters for the password. Default is 128.

- `-c, --config string`  
  **Path to a configuration file**  
  Specify the path to a configuration file if needed (currently a placeholder for future enhancements).

- `-v, --verbose`  
  **Verbose output**  
  Enable verbose mode to display detailed information about the password generation process.


## Examples
Generate a 12-character password with symbols and numbers:

```bash
./password_generator --length 12 --symbols --numbers
```
Generate a 20-character password with verbose output and excluding similar characters:

```bash
./password_generator --length 20 --exclude-similar --verbose
```
Generate a password between 8 and 16 characters long:

```bash
./password_generator --min-length 8 --max-length 16
```

## Contributing
Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

Fork the Repository: Click the "Fork" button at the top right of this page to create your own copy of the repository.

Clone Your Fork:

```bash
git clone https://github.com/JakubSchwenkbeck/password-generator.git
```
Create a Branch:

```bash
git checkout -b my-feature-branch
```
Make Changes: Implement your feature or fix a bug.

Commit Changes:

```bash
git add .
git commit -m "Add feature or fix bug"
```

Push Changes:

```bash
git push origin my-feature-branch
```
Open a Pull Request: Go to the original repository and create a pull request from your fork.

Please ensure that your code follows the project's style and includes tests where applicable.

License
This project is licensed under the MIT License - see the LICENSE file for details.

Acknowledgments
Thanks to the Go community for the powerful libraries and tools.
Special thanks to the maintainers of the cobra library for building a great CLI framework.
Contact
For any questions or feedback, please open an issue on the GitHub repository.
