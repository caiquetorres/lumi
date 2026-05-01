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
    let n = 10
    println(fibonacci(n))
}

fun fibonacci(n int) {
    let a = 0
    let b = 1

    let i = 0
    while i < n {
        let tmp = a + b
        b = a
        a = tmp

        i = i + 1
    }

    return a
}
```
