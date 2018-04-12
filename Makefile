PKGS := $(shell go list ./... | grep -v /vendor)

.default: test

.PHONY: test
test: lint
	go test $(PKGS)

BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: dependencies
dependencies:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

.PHONY: lint
lint: $(GOMETALINTER) dependencies
	gometalinter ./... --vendor -e "Subprocess launching with variable"

BINARY := clp
PLATFORMS := windows linux darwin
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release/
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(os)

.PHONY: release
release: windows linux darwin

.PHONY: clean
clean:
	rm -r release/

.PHONY: all
all: test release
