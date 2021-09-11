generate:
	pigeon grammar.peg | gofmt > grammar.go

install:
	go install github.com/mna/pigeon@latest

bench:
	go test -benchmem -bench=.

test:
	go test -v -cover .
