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
    let firstName = "Jane";
    let lastName = "Doe";

    let name = getPreferredName(firstName, lastName);
    let greeting = buildGreeting(name);

    printf("%s, %s!\n", greeting, name);
}

fun getPreferredName(first string, last string) string {
    return first;
}

fun buildGreeting(name string) string {
    return "Hello";
}
```
