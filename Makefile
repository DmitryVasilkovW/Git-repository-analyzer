VERSION=v1.0.0
BINARY=gitanalyzer
FILES=$(BINARY) README.md

GOOS=darwin
GOARCH=amd64

build-zip-mac:
	zip -r $(BINARY)-$(VERSION)-$(GOOS)-$(GOARCH).zip $(FILES)

build-tar-mac:
	tar -czvf $(BINARY)-$(VERSION)-$(GOOS)-$(GOARCH).tar.gz $(FILES)

GOOS=windows
GOARCH=amd64

build-zip-win:
	zip -r $(BINARY)-$(VERSION)-$(GOOS)-$(GOARCH).zip $(BINARY).exe $(FILES)

build-tar-win:
	tar -czvf $(BINARY)-$(VERSION)-$(GOOS)-$(GOARCH).tar.gz $(BINARY).exe $(FILES)

GOOS=linux
GOARCH=amd64

build-zip-linux:
	zip -r $(BINARY)-$(VERSION)-$(GOOS)-$(GOARCH).zip $(FILES)

build-tar-linux:
	tar -czvf $(BINARY)-$(VERSION)-$(GOOS)-$(GOARCH).tar.gz $(FILES)

build:
	go build -o gitanalyzer ./cmd/analyzer

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o gitanalyzer ./cmd/analyzer

build-win:
	GOOS=windows GOARCH=amd64 go build -o gitanalyzer.exe ./cmd/analyzer

build-linux:
	GOOS=linux GOARCH=amd64 go build -o gitanalyzer ./cmd/analyzer
