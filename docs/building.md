# How to build

### Requirements

- [Golang](https://golang.org) installed on your system
- A C compiler (for cgo, docker images compiled with gcc and musl libc)

### Steps to build

- Clone this repository
- Run `go build -o passwdbox -ldflags "-s" cmd/passwdbox/main.go` to build the binary
- The `passwdbox` binary will be available in your current directory
