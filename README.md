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
    for i in 2..=20 {
        if isPrime(i) {
            printf("%d is prime\n", i)
        }
    }
}

fun isPrime(n int) {
    if n < 2 {
        return false
    }
    let i = 2
    while i * i <= n {
        if n / i * i == n {
            return false
        }
        i += 1
    }
    return true
}
```
