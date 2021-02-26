TARGETS := amd64 arm64 win64
export ARCH_DATALAYER_arm64 := linux-gcc-aarch64
export ARCH_DATALAYER_amd64 := linux-gcc-x64
export CROSS_COMPILE_arm64 := aarch64-linux-gnu-
export CGO_CFLAGS=-I${PWD}/deps/public/include/comm.datalayer/comm/datalayer/c
export CGO_LDFLAGS_linux="-L${PWD}/deps/public/bin/comm.datalayer/$${ARCH_DATALAYER_$*}/release -L${PWD}/deps/public/bin/zmq/$${ARCH_DATALAYER_$*}/release"

.PHONY: build
build: $(patsubst %,build_%,$(TARGETS))

.PHONY: build_linux
build_linux: $(patsubst %,build_%,amd64 arm64)

.PHONY: build_%
build_%:
	mkdir -p build/public/linux_$*
	for x in ./cmd/*; do \
		CGO_LDFLAGS=$(CGO_LDFLAGS_linux)\
		GOARCH=$* GOOS=linux CGO_ENABLED=1 CC=$${CROSS_COMPILE_$*}gcc go build -mod vendor -o build/public/linux_$*/$$(basename $$x) $$x; \
	done

.PHONY: build_win64
build_win64:
	mkdir -p build/public/win_amd64
	mkdir -p build/temp/win64
	cp deps/public/bin/comm.datalayer/win-msvc-x64/release/comm_datalayer.dll build/temp/win64
	for x in ./cmd/*; do \
		CGO_LDFLAGS="-L${PWD}/build/temp/win64" \
		GOARCH=amd64 GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc \
		go build -mod vendor -o build/public/win_amd64/$$(basename $$x).exe $$x; \
	done

.PHONY: test
test: test_amd64

.PHONY: test_%
test_%:
	LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:${PWD}/deps/public/bin/comm.datalayer/$${ARCH_DATALAYER_$*}/release:${PWD}/deps/public/bin/zmq/linux-gcc-x64/release \
	CGO_LDFLAGS=$(CGO_LDFLAGS_linux) \
	go test ./...
	