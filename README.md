# godoc-examples

Dummy project with testable examples that get embedded into docs when using `godoc`

### Links

- Official go blog article: https://go.dev/blog/examples
- Go docs: https://pkg.go.dev/testing#hdr-Examples
- `godoc` package https://pkg.go.dev/golang.org/x/tools/cmd/godoc

### Install `godoc`

```bash
$ go install golang.org/x/tools/cmd/godoc@latest
```

### Run

```bash
$ godoc -http=:6060
```

### Run tests

```bash
$ go test ./...
```
