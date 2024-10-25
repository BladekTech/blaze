.PHONY: build
.PHONY: run
.PHONY: package
.PHONY: cli
.PHONY: server

package: build docs examples
	tar -cvzf dist 

dist/examples:
	mkdir -p dist/docs
	cp examples dist/examples

dist/docs:
	mkdir -p dist/docs

build: clean dist/bin/blaze-server.exe dist/bin/blaze-cli.exe

clean:
	rm -f dist/bin/*

run-server: clean-server dist/bin/blaze-server.exe
	./dist/bin/blaze-server.exe

run-cli: clean-cli dist/bin/blaze-cli.exe
	./dist/bin/blaze-cli.exe

dist/bin:
	mkdir -p dist/bin

cli: clean-cli dist/bin/blaze-cli.exe
server: clean-server dist/bin/blaze-server.exe

dist/bin/blaze-server.exe: dist/bin
	go build ./cmd/blaze-server
	mv ./blaze-server.exe ./dist/bin/blaze-server.exe

dist/bin/blaze-cli.exe: dist/bin
	go build ./cmd/blaze-cli
	mv ./blaze-cli.exe ./dist/bin/blaze-cli.exe

clean-cli:
	rm -f dist/bin/blaze-cli.exe

clean-server:
	rm -f dist/bin/blaze-server.exe
