PKGS := $(shell go list ./... | grep -v /vendor)

.default: test


BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter
GOCOV := $(BIN_DIR)/gocov
GOCOV_XML := $(BIN_DIR)/gocov-xml
GO_JUNIT_REPORT := $(BIN_DIR)/go-junit-report

.PHONY: test
test: lint $(GOCOV) $(GOCOV_XML) $(GO_JUNIT_REPORT) reports_prepare
	go test -v $(PKGS)
	go test -v $(PKGS) | $(GO_JUNIT_REPORT) > reports/test.xml
	$(GOCOV) test $(PKGS) | $(GOCOV_XML) > reports/coverage.xml

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

$(GOCOV):
	go get -u github.com/axw/gocov/...

$(GOCOV_XML):
	go get -u github.com/AlekSi/gocov-xml

$(GO_JUNIT_REPORT):
	go get -u github.com/jstemmer/go-junit-report

.PHONY: dependencies
dependencies:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

.PHONY: reports_prepare
reports_prepare:
	mkdir -p reports

.PHONY: lint
lint: $(GOMETALINTER) dependencies reports_prepare
	$(GOMETALINTER) ./... --vendor -e "Subprocess launching with variable"
	$(GOMETALINTER) ./... --vendor -e "Subprocess launching with variable" --checkstyle > reports/report.xml

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

.PHONY: release_tag
release_tag:
	git tag `date +%Y%m%d%H%M`
