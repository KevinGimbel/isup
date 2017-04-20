# isup

[![Go Report Card](https://goreportcard.com/badge/github.com/kevingimbel/isup.go)](https://goreportcard.com/report/github.com/kevingimbel/isup.go)

Go executable to check if a remote host is available or not. First argument passed is the URL to check. **This is a test and my first test of programming in Go beside Hello World.**

### Install
1. [Install Go](https://golang.org/doc/install#install).
2. Clone the repo `git clone https://github.com/kevingimbel/isup.go.git`
3. Run `go build isup.go` from within the new directory.

### Usage
`isup` can be used from the command line as follows:

```sh
$ ./path/to/isup "http://kevingimbel.com"
```
If you want to use it globally, create a symlink to the executable somewhere inside your`$PATH`.
