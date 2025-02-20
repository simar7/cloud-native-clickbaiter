all: build

build: binary

binary: $(shell find . -name '*.go')
	CGO_ENABLED=0 go build -o app -v -ldflags '-w -extldflags '-static''

container:
	docker build -t quay.io/cbrgm/clickbaiter:latest .