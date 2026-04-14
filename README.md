# Lumi

A simple, statically-typed programming language written in Go. Lumi compiles to bytecode and runs on a custom virtual machine.

```sh
# Build a Lumi source file to bytecode
go run . build -file <file>.lumi -out <file>.bin

# Run a compiled bytecode file
go run . run <file>.bin
```

## Example

```lumi
fun main() {
    sayHelloTo("Jane")
}

fun sayHelloTo(name string) {
    printf("Hello %s how are you?\n", name)
}
```
