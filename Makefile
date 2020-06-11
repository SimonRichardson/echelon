.PHONY: setup, test, unit-test, test-ci

project=echelon

TARGET ?= echelon
TAG ?= dev

setup:
	go get -u github.com/jteeuwen/go-bindata/...
	make build-bindata

test:
	make TARGET=echelon test-target
	make TARGET=echelonw test-target
	make TARGET=echelons test-target

.PHONY: build-bindata, build-schemas

build-bindata:
	cd scripts && go-bindata --nocompress -pkg scripts ./../scripts/... && cd ../

build-schemas:
	@rm -rf ./schemas/schema
	flatc -g -o ./schemas/ ./schemas/*fbs

.PHONY: build, build-image, build-images, publish-images

build:
	mkdir -p artifacts
	CGO_ENABLED=0 go build -o bin/echelon github.com/SimonRichardson/echelon/echelon-http
	tar -czf artifacts/echelon.tar.gz bin/echelon -C bin/ .
	CGO_ENABLED=0 go build -o bin/echelonw github.com/SimonRichardson/echelon/echelon-walker
	tar -czf artifacts/echelonw.tar.gz bin/echelonw -C bin/ .
	CGO_ENABLED=0 go build -o bin/echelons github.com/SimonRichardson/echelon/echelon-shim
	tar -czf artifacts/echelons.tar.gz bin/echelons -C bin/ .
.PHONY: internal-echelon-build, internal-echelonw-build, internal-echelons-build

internal-echelon-build:
	@make setup -s
	CGO_ENABLED=0 go build -o bin/echelon github.com/SimonRichardson/echelon/echelon-http
	tar -czf - echelon-http/Dockerfile bin/echelon

internal-echelonw-build:
	@make setup -s
	CGO_ENABLED=0 go build -o bin/echelon github.com/SimonRichardson/echelon/echelon-walker
	tar -czf - echelon-walker/Dockerfile bin/echelonw

internal-echelons-build:
	@make setup -s
	CGO_ENABLED=0 go build -o bin/echelon github.com/SimonRichardson/echelon/echelon-shim
	tar -czf - echelon-shim/Dockerfile bin/echelons

.PHONY: internal-echelon-tests, internal-echelon-integration-tests, internal-echelon-benchmark-tests, internal-echelon-perf-tests

## echelon

internal-echelon-tests:
	go test -v ./cluster/counter/... -stubs=true
	go test -v ./cluster/counter/...
	go test -v ./cluster/store/... -stubs=true
	go test -v ./cluster/store/...
	go test -v ./selectors/...
	go test -v ./semaphore/...

internal-echelon-integration-tests:
	go test -v ./echelon-http/...

internal-echelon-benchmark-tests:
	go test -bench=. -run=BenchmarkRequest ./echelon-http/...

internal-echelon-perf-tests:
	sh -c 'go run echelon-perf/*.go -ciserver=true -clusterserver=true'

.PHONY: internal-echelonw-tests, internal-echelonw-integration-tests, internal-echelonw-benchmark-tests, internal-echelonw-perf-tests

## Walker

internal-echelonw-tests:
	# Do nothing!

internal-echelonw-integration-tests:
	go test -v ./echelon-walker/...

internal-echelonw-benchmark-tests:
	# Do nothing!

internal-echelonw-perf-tests:
	# Do nothing!

.PHONY: internal-echelons-tests, internal-echelons-integration-tests, internal-echelons-benchmark-tests, internal-echelons-perf-tests

## Shim

internal-echelons-tests:
	# Do nothing!

internal-echelons-integration-tests:
	CGO_ENABLED=0 go build -o bin/echelon github.com/SimonRichardson/echelon/echelon-http
	./bin/echelon &
	sleep 4
	go test -v ./echelon-shim/...

internal-echelons-benchmark-tests:
	# Do nothing!

internal-echelons-perf-tests:
	# Do nothing!
