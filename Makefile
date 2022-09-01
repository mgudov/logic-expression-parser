generate:
	pigeon grammar.peg | gofmt > grammar.go

install:
	go install github.com/mna/pigeon@latest

cover:
	go test ./... -coverprofile cover.out
	go tool cover -func cover.out | grep -v grammar.go

bench:
	go test -benchmem -bench=.

test:
	go test -v -cover .
