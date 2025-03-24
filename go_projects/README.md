# Go programming language

## History of Go

### Why was Go created?

* Long compilation times
* Slow builds
* Uncontrolled dependencies
* Unreadable code
* Poor concurrency support
* Poor dependency management
* Poor error handling

### Who created Go?

Go was created at Google by:

* Robert Griesemer - (worked on V8, HotSpot, and Strongtalk at Google)
* Rob Pike - (wrote book "The Practice of Programming" with Brian Kernighan, worked on Unix tools, Plan 9, UTF-8)
* Ken Thompson  - (also involved in the creation of Unix, B and C, grep, ed, Plan 9, UTF-8, and the first compiler for the C programming language)

### When was Go created?

Go was created in 2007, but it was not released to the public until 2009.

#### Timeline of Go releases

From: https://go.dev/doc/devel/release

* 2009 - Pre 1.0
* 2012 - Go 1
* 2013 - Go 1.1
* 2013 - Go 1.2

* 2025 - Go 1.25

## Go's design goals

* Simplicity
* Fast compilation
* Garbage collection
* Concurrency
* Static typing
* Memory safety
* Strong tooling
* Cross-compilation
* Standard library
* Dependency management
* Error handling
* Testing
* Profiling
* Documentation
* Code formatting

## Go's usage

Go is used for:

* Web servers
* Command-line tools
* Desktop applications to some extent
* Cloud services
* Networking
* Databases
* Machine learning a little bit
* Embedded systems
* Mobile applications a little bit
* WebAssembly
* Games a little bit
* Audio and video processing a little bit
* Cryptography
* Blockchain
* DevOps
* Testing
* Automation
* Data processing

## Running and Building Go programs

### Running Go programs

To run a Go program, you can use the `go run` command followed by the name of the file you want to run.

```bash
go run main.go
```

### Building Go programs

To build a Go program, you can use the `go build` command followed by the name of the file you want to build.

```bash
go build main.go
```
 This will build a binary file with the same name as the file you built in this case `main`.

 This binary can be run on any machine that has the same computer architecture as the one you built it on.