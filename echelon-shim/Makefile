GO ?= go

all: build

setup:

build:
	$(GO) build

clean:
	$(GO) clean

check:
	@$(GO) list -f '{{join .Deps "\n"}}' | xargs $(GO) list -f '{{if not .Standard}}{{.ImportPath}} {{.Dir}}{{end}}' | column -t
