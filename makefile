.DEFAULT_GOAL := build

fmt:
	gofmt -s -w .
.PHONY:fmt

build: fmt
	go build -o dist/gmdump main.go
	./dist/gmdump
.PHONY:build

debug:
	go build -gcflags '-N -l' -o dist/debug main.go
.PHONY:debug
