APP=gbot
REGISTRY=stas727
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux
TARGETARCH=amd64

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

goget:
	go get

build: format
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${LINUX_TARGETARCH} GOENV=.env go build -v -o gbot -ldflags "-X="github.com/stas727/gbot/cmd.appVersion=${VERSION}

image:
	docker build  --platform ${TARGETOS}/${TARGETARCH} . -t ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

push:
	docker push ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

clean:
	rm -rf gbot && docker image rm ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}
