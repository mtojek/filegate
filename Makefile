build: clean-code install test

clean-code:
	go get golang.org/x/tools/cmd/goimports && goimports -w .
	gofmt -s -w .
	go get golang.org/x/lint/golint && golint -set_exit_status ./...

install:
	go get -v github.com/mtojek/filegate/cmd/filegate

test: install
	go get -t ./...
	go test -v ./...
	filegate || test -n "$$?"
	filegate version
	filegate pull --help
	filegate push --help
	filegate signal --help
