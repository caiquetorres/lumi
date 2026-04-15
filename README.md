# Lumi

> **Work in Progress** — This language is not complete. Features and syntax are subject to change.

A simple, statically-typed programming language written in Go. Lumi compiles to bytecode and runs on a custom virtual machine.

```sh
# Build a Lumi source file to bytecode
go run . build -file <file>.lumi -out <file>.lbc

# Run a compiled bytecode file
go run . run <file>.lbc
```

## Example

```lumi
fun main() {
    let first = "John";
    let last = "Doe";
    introduce(first, last);
}

fun introduce(first string, last string) {
    printf("My name is %s %s\n", first, last);
}
```
