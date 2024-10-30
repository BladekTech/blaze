.PHONY: build
.PHONY: run
.PHONY: package
.PHONY: cli
.PHONY: server
.PHONY: ucache
.PHONY: build-server
.PHONY: build-server-windows
.PHONY: build-server-linux
.PHONY: build-server-darwin

package: build examples
	tar -cvzf blaze.tar.gz ./dist

dist/examples:
	cp examples dist/examples
	
build: build-server dist/bin/blaze-windows-setup.exe dist/bin/blaze-cli.exe

clean: clean-server clean-cli

dist/bin:
	mkdir -p dist/bin
dist/bin/windows: dist/bin
	mkdir -p dist/bin/windows
dist/bin/linux: dist/bin
	mkdir -p dist/bin/linux
dist/bin/darwin: dist/bin
	mkdir -p dist/bin/darwin

cli: clean-cli dist/bin/blaze-cli.exe

build-server: build-server-windows build-server-linux build-server-darwin

# windows
build-server-windows: dist/bin/windows/blaze-server-windows-386.exe dist/bin/windows/blaze-server-windows-amd64.exe dist/bin/windows/blaze-server-windows-arm.exe dist/bin/windows/blaze-server-windows-arm64.exe

dist/bin/windows/blaze-server-windows-386.exe: clean-server-windows-386 dist/bin/windows
	GOOS=windows GOARCH=386 \
	go build ./cmd/blaze-server
	mv ./blaze-server.exe ./dist/bin/windows/blaze-server-windows-386.exe
dist/bin/windows/blaze-server-windows-amd64.exe: clean-server-windows-amd64 dist/bin/windows
	GOOS=windows GOARCH=amd64 \
	go build ./cmd/blaze-server
	mv ./blaze-server.exe ./dist/bin/windows/blaze-server-windows-amd64.exe
dist/bin/windows/blaze-server-windows-arm.exe: clean-server-windows-arm dist/bin/windows
	GOOS=windows GOARCH=arm \
	go build ./cmd/blaze-server
	mv ./blaze-server.exe ./dist/bin/windows/blaze-server-windows-arm.exe
dist/bin/windows/blaze-server-windows-arm64.exe: clean-server-windows-arm64 dist/bin/windows
	GOOS=windows GOARCH=arm64 \
	go build ./cmd/blaze-server
	mv ./blaze-server.exe ./dist/bin/windows/blaze-server-windows-arm64.exe

dist/bin/blaze-windows-setup.exe: dist/bin
	rm -f dist/bin/blaze-window-setup.exe
	iscc blaze-setup-win.iss

# linux
build-server-linux: dist/bin/linux/blaze-server-linux-386.exe dist/bin/linux/blaze-server-linux-amd64.exe dist/bin/linux/blaze-server-linux-arm.exe dist/bin/linux/blaze-server-linux-arm64.exe

dist/bin/linux/blaze-server-linux-386.exe: clean-server-linux-386 dist/bin/linux
	GOOS=linux GOARCH=386 \
	go build ./cmd/blaze-server
	mv ./blaze-server ./dist/bin/linux/blaze-server-linux-386.exe
dist/bin/linux/blaze-server-linux-amd64.exe: clean-server-linux-amd64 dist/bin/linux
	GOOS=linux GOARCH=amd64 \
	go build ./cmd/blaze-server
	mv ./blaze-server ./dist/bin/linux/blaze-server-linux-amd64.exe
dist/bin/linux/blaze-server-linux-arm.exe: clean-server-linux-arm dist/bin/linux
	GOOS=linux GOARCH=arm \
	go build ./cmd/blaze-server
	mv ./blaze-server ./dist/bin/linux/blaze-server-linux-arm.exe
dist/bin/linux/blaze-server-linux-arm64.exe: clean-server-linux-arm64 dist/bin/linux
	GOOS=linux GOARCH=arm64 \
	go build ./cmd/blaze-server
	mv ./blaze-server ./dist/bin/linux/blaze-server-linux-arm64.exe

# darwin
build-server-darwin: dist/bin/darwin/blaze-server-darwin-amd64.exe dist/bin/darwin/blaze-server-darwin-arm64.exe

dist/bin/darwin/blaze-server-darwin-amd64.exe: clean-server-darwin-amd64 dist/bin/darwin
	GOOS=darwin GOARCH=amd64 \
	go build ./cmd/blaze-server
	mv ./blaze-server ./dist/bin/darwin/blaze-server-darwin-amd.exe
dist/bin/darwin/blaze-server-darwin-arm64.exe: clean-server-darwin-arm64 dist/bin/darwin
	GOOS=darwin GOARCH=arm64 \
	go build ./cmd/blaze-server
	mv ./blaze-server ./dist/bin/darwin/blaze-server-darwin-arm64.exe

dist/bin/blaze-cli.exe: dist/bin
	go build ./cmd/blaze-cli
	mv ./blaze-cli.exe ./dist/bin/blaze-cli.exe

clean-cli:
	rm -f dist/bin/blaze-cli.exe

clean-server: clean-server-windows clean-server-linux clean-server-darwin
clean-server-windows:
	rm -f dist/bin/windows/blaze-server-windows-*.exe
clean-server-linux:
	rm -f dist/bin/linux/blaze-server-linux-*.exe
clean-server-darwin:
	rm -f dist/bin/darwin/blaze-server-darwin-*.exe

clean-server-windows-386:
	rm -f dist/bin/windows/blaze-server-windows-386.exe
clean-server-windows-amd64:
	rm -f dist/bin/windows/blaze-server-windows-amd64.exe
clean-server-windows-arm64:
	rm -f dist/bin/windows/blaze-server-windows-arm64.exe
clean-server-windows-arm:
	rm -f dist/bin/windows/blaze-server-windows-arm.exe

clean-server-linux-386:
	rm -f dist/bin/linux/blaze-server-linux-386.exe
clean-server-linux-amd64:
	rm -f dist/bin/linux/blaze-server-linux-amd64.exe
clean-server-linux-arm64:
	rm -f dist/bin/linux/blaze-server-linux-arm64.exe
clean-server-linux-arm:
	rm -f dist/bin/linux/blaze-server-linux-arm.exe

clean-server-darwin-amd64:
	rm -f dist/bin/darwin/blaze-server-darwin-amd64.exe
clean-server-darwin-arm64:
	rm -f dist/bin/darwin/blaze-server-darwin-arm64.exe


# Updates Proxy Cache
ucache:
	echo "Updating proxy cache to version $(VERSION)"
    GOPROXY=https://proxy.golang.org GO111MODULE=on \
    go get github.com/BladekTech/blaze@v$(VERSION)
