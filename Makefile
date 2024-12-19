VERSION=v1.0.0

# archive

# mac

MAC_GOOS=mac
MAC_GOARCH=amd64
MAC_BINARY=gitanalyzer
MAC_FILES=$(MAC_BINARY) README.md

build-zip-mac:
	zip -r $(MAC_BINARY)-$(VERSION)-$(MAC_GOOS)-$(MAC_GOARCH).zip $(MAC_FILES)

build-tar-mac:
	tar -czvf $(MAC_BINARY)-$(VERSION)-$(MAC_GOOS)-$(MAC_GOARCH).tar.gz $(MAC_FILES)

# windows

WINDOWS_GOOS=windows
WINDOWS_GOARCH=amd64
WINDOWS_BINARY=gitanalyzer.exe
WINDOWS_FILES=$(WINDOWS_BINARY) README.md

build-zip-win:
	zip -r $(WINDOWS_BINARY)-$(VERSION)-$(WINDOWS_GOOS)-$(WINDOWS_GOARCH).zip $(WINDOWS_FILES)

build-tar-win:
	tar -czvf $(WINDOWS_BINARY)-$(VERSION)-$(WINDOWS_GOOS)-$(WINDOWS_GOARCH).tar.gz $(WINDOWS_FILES)

# linux

LINUX_GOOS=linux
LINUX_GOARCH=amd64
LINUX_BINARY=gitanalyzer
LINUX_FILES=$(LINUX_BINARY) README.md

build-zip-linux:
	zip -r $(LINUX_BINARY)-$(VERSION)-$(LINUX_GOOS)-$(LINUX_GOARCH).zip $(LINUX_FILES)

build-tar-linux:
	tar -czvf $(LINUX_BINARY)-$(VERSION)-$(LINUX_GOOS)-$(LINUX_GOARCH).tar.gz $(LINUX_FILES)


# build

build:
	go build -o gitanalyzer ./cmd/analyzer

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o gitanalyzer ./cmd/analyzer

build-win:
	GOOS=windows GOARCH=amd64 go build -o gitanalyzer.exe ./cmd/analyzer

build-linux:
	GOOS=linux GOARCH=amd64 go build -o gitanalyzer ./cmd/analyzer


get-mac-release: build-mac build-zip-mac build-tar-mac
get-win-release: build-win build-zip-win build-tar-win
get-linux-release: build-linux build-zip-linux build-tar-linux