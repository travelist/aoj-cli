NAME            := aoj
VERSION         := v0.5.0
REVISION        := $(shell git rev-parse --short HEAD)
LDFLAGS         := "-X github.com/travelist/aoj-cli/cmd.Version=${VERSION} -X github.com/travelist/aoj-cli/cmd.Revision=${REVISION}"

.PHONY: build
build:
	go build -ldflags $(LDFLAGS) -o aoj

