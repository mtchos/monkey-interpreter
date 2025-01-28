# Monkey Interpreter

This project implements the **Monkey programming language**, following along with the book:

**[Writing An Interpreter In Go](https://interpreterbook.com/)** by Thorsten Ball.

Monkey is a simple, toy programming language designed to teach the fundamental concepts of building an interpreter. This project demonstrates key processes such as lexing, parsing, evaluating, and extending a programming language, all implemented in Go.

### Features
- **Lexing**: Tokenizes source code into meaningful units.
- **Parsing**: Builds an Abstract Syntax Tree (AST) from tokens.
- **Evaluating**: Executes Monkey code based on the AST.
- **Extensible Data Structures**: Includes support for integers, booleans, and strings.
- **Built-in Functions**: Includes functions like `print` for debugging and displaying values.
- Simple error handling to highlight invalid code.

### Learning Objectives
- Understand the building blocks of an interpreter.
- Explore how lexing, parsing, and evaluation work together.
- Learn how to extend a language with built-in functions and custom features.

### Installation

1. Clone the repository:
   ```bash
   git clone git@github.com:mtchos/monkey-interpreter.git
   ```
2. Navigate to the repository folder:
   ```bash
   cd monkey-interpreter
   ```
3. Build the project:
   ```bash
   go build
   ```
4. Execute the project:
   ```bash
   ./monkey-interpreter
   ```
5. Press `Ctrl + C` to exit the interpreter.

### Example Usage

Run the interpreter and try out the Monkey programming language:

```bash
$ ./monkey-interpreter
>> let five = 5;
>> print(five);
5
>> let ten = 10;
>> print(ten);
10
>> print(five + ten);
15
```

### Contributing
Feel free to submit issues and pull requests if you discover bugs or have ideas for improving the interpreter.

### License
This project is licensed under the MIT License. See the [LICENSE](https://mit-license.org/) file for more details.
