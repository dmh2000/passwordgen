# Password Generator CLI

A command-line utility for generating secure passwords with customizable options.

## Features

- Generates passwords with a minimum length of 24 characters
- Option to include special symbols
- Supports generating multiple passwords at once
- Outputs both raw and formatted password versions
- Provides both short and long-form command flags

## Usage

```bash
password-generator [-l length] [-c count] [-s] [-h]
```

### Flags

- `-l, --length`: Set password length (minimum 24 characters, default: 24)
- `-c, --count`: Number of passwords to generate (default: 1)
- `-s, --symbols`: Include special symbols in the password
- `-h, --help`: Display help message

## Output

For each password generated, the program outputs:
- A separator line
- The raw password
- A formatted version of the password

## Example

```bash
password-generator -l 30 -c 2 -s
```

This command generates 2 passwords, each 30 characters long, including special symbols.

## Notes

- The program enforces a minimum password length of 24 characters for security
- Both short (-l) and long (--length) format flags are supported
- If an error occurs during generation or formatting, the program will exit with an error message

The implementation uses a separate package for password generation and formatting logic, keeping the CLI interface focused on user interaction and parameter handling.
