APP=gbot
REGISTRY=gcr.io/gbot-418317
VERSION=$(shell echo 'v1.0.0')-$(shell git rev-parse --short HEAD)

LINUX_TARGETOS=linux
LINUX_TARGETARCH=arm64
ARM_TARGETOS=android
ARM_TARGETARCH=arm64
MACOS_TARGETOS=darwin
MACOS_TARGETARCH=amd64
WINDOWS_TARGETOS=windows
WINDOWS_TARGETARCH=arm64
#AVAILABLE GOOS_VALUES=["aix", "android", "darwin", "dragonfly", "freebsd", "hurd", "illumos", "ios", "js", "linux", "nacl", "netbsd", "openbsd", "plan9", "solaris", "windows", "zos"]
#AVAILABLE GOARCH_VALUES=["386", "amd64", "amd64p32", "arm", "arm64", "arm64be", "armbe", "loong64", "mips", "mips64", "mips64le", "mips64p32", "mips64p32le", "mipsle", "ppc", "ppc64", "ppc64le", "riscv", "riscv64", "s390", "s390x", "sparc", "sparc64", "wasm"]
format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

goget:
	go get

linux: format
	CGO_ENABLED=0 GOOS=${LINUX_TARGETOS} GOARCH=${LINUX_TARGETARCH} go build -v -o gbot -ldflags "-X="github.com/stas727/gbot/cmd.appVersion=${VERSION}

arm: format
	CGO_ENABLED=0 GOOS=${ARM_TARGETOS} GOARCH=${ARM_TARGETARCH} go build -v -o gbot -ldflags "-X="github.com/stas727/gbot/cmd.appVersion=${VERSION}

macos: format
	CGO_ENABLED=0 GOOS=${MACOS_TARGETOS} GOARCH=${MACOS_TARGETARCH} go build -v -o gbot -ldflags "-X="github.com/stas727/gbot/cmd.appVersion=${VERSION}

windows: format
	CGO_ENABLED=0 GOOS=${WINDOWS_TARGETOS} GOARCH=${WINDOWS_TARGETARCH} go build -v -o gbot -ldflags "-X="github.com/stas727/gbot/cmd.appVersion=${VERSION}

image:
	docker build . -t ${REGISTRY}/${APP}:${VERSION}

push:
	docker push ${REGISTRY}/${APP}:${VERSION}

clean:
	rm -rf gbot && docker image rm ${REGISTRY}/${APP}:${VERSION}
