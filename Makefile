VERSION=v1.0.0
BINARY=gitanalyzer
FILES=$(BINARY) README.md
GOOS=darwin
GOARCH=amd64

build-zip:
	zip -r $(BINARY)-$(VERSION)-$(GOOS)-$(GOARCH).zip $(FILES)

build-tar:
	tar -czvf $(BINARY)-$(VERSION)-$(GOOS)-$(GOARCH).tar.gz $(FILES)

build:
	go build -o gitanalyzer ./cmd/analyzer
