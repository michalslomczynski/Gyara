# About
Gyara is a script and library that lexes and parses a file consisting of one more YARA rules into a Golang struct type representation. The goal of this tool is to import and parse the contents of multiple YARA rules directly within Go code.

# How to use
Import and call ParseRule in your codebase.\
Execute locally\
```bash

go build -o gyara cli/main.go && ./gyara <rule_file>
```
To install (linux)
```bash
sudo make
```
