.PHONY: all clean
TARGETNAME=ui-api
GOOS=linux
GOARCH=amd64

all: format test build clean

test:
	go test -v . 

format:
	gofmt -w .

build:
	mkdir -p releases
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 CGO_LDFLAGS="-static" go build -mod=vendor -ldflags "-s -w" -v -o releases/$(TARGETNAME) .

clean:
	go clean -i
