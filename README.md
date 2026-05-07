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
    for let n = 1; n <= 15; n += 1 {
        let steps = collatz(n)
        printf("collatz(%d): %d steps\n", n, steps)
    }
}

fun collatz(n int) int {
    let steps = 0
    while n != 1 {
        if n / 2 * 2 == n {
            n /= 2
        } else {
            n = 3 * n + 1
        }
        steps += 1
    }
    return steps
}
```
