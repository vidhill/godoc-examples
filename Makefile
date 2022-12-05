SHELL := /bin/bash # mac quirk, need to declare which shell to use
GOPATH := $(shell go env GOPATH)

test:
		go test ./...

serve-docs: check.godoc
		godoc -http=:6060

check.godoc:
   ifeq (, $(shell which ${GOPATH}/bin/godoc))
		$(error godoc is not installed, Please install by runing `go install golang.org/x/tools/cmd/godoc@latest`)
   endif