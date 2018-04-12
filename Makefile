PKGS := $(shell go list ./... | grep -v /vendor)

.default: test

.PHONY: test
test: lint
	go test $(PKGS)

BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

.PHONY: lint
lint: $(GOMETALINTER)
	gometalinter ./... --vendor -e "Subprocess launching with variable"

BINARY := clp
PLATFORMS := windows linux darwin
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release/$(os)
	GOOS=$(os) GOARCH=amd64 go build -o release/$(os)/$(BINARY)

.PHONY: release
release: windows linux darwin

.PHONY: clean
clean:
	rm -r release/

.PHONY: all
all: test release