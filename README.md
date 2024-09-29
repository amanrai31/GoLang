# GoLang
Go language

Go lang is multithreaded thats why it has some disadvantages too like concurrency.

- Built-in concurrency machanism = `c++`,`java`,`GO`

- No Built-in concurrency machanism = python, node.

***Note -*** But `c++`and `java` have complex code(human errors), expensive & slow. Go runs on multiple cores and supports concurrency with ease.

#### Main use cases of Go
1. Made for performant application.
2. Runs on scaled, distributed systems.

#### Characterstics of Go (Simplicity and speed)
1. Simple and readable syntax of dynamic typed  langs(more hifh level lang) like python.
2. Speed,safty and efficiency like lower level/statically typed langs like c++/java.
3. It is complied lang, compiles into single binary. (Take that binary and use across diff plateforms.)
4. Consistent across diff OS.

---

1. go mod init <project-name>
2. Create main.go file
3. `package main` in file `main.go`- All code belong to a package.
4. `func main()` is entry point in Go. A program(app) can have only 1 main function
5. `import "fmt"` core package provided by GO library.

- ***Note -*** Go prorams are organized into pakages, Go's standard library provides diff core packages for us- `fmt` is one of them.

6. `go run <filename>` e.g.- go run main.go