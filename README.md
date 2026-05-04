# Lumi

> **Work in Progress** — This language is not complete. Features and syntax are subject to change.

A simple, statically-typed programming language written in Go. Lumi compiles to bytecode and runs on a custom virtual machine.

```sh
# Build a Lumi source file to bytecode
go run . build -file <file>.lumi -out <file>.lbc

# Run a compiled bytecode file
go run . run <file>.lbc
```

## Examples

More examples can be found in the [`examples/`](examples/) folder.

```lumi
fun main() {
    for let i = 1; i <= 10; i += 1 {
        let fib = fibonacci(i)
        printf("fib(%d): %d\n", i, fib)
    }
}

fun fibonacci(n int) {
    let a = 0, b = 1

    for let i = 0; i < n; i += 1 {
        let tmp = a + b
        b = a
        a = tmp
    }

    return a
}
```
