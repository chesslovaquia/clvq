# host targets

.PHONY: docker
docker:
	./docker/build.sh

# container targets

.PHONY: fmt
fmt:
	@gofmt -l -s -w .

.PHONY: all
all: build

.PHONY: build
build:
	go install ./cmd/clvq

.PHONY: upgrade
upgrade:
	@go version
	go get go@latest
	go get -u all
	go mod tidy -v
#	go mod vendor
