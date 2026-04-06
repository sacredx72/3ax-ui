BINARY   := 3ax-ui
TARGET   := target
TMP      := tmp
MODULE   := github.com/coinman-dev/3ax-ui/v2/config

# Version: prefer git tag, fallback to config/version file
VERSION  := $(shell git describe --tags --always --dirty 2>/dev/null || cat config/version)
LDFLAGS  := -X '$(MODULE).version=$(VERSION)'

.PHONY: all build clean tmp-dir target-dir

all: build

## build — compile binary into target/ with version from git tag
build: target-dir
	go build -ldflags "$(LDFLAGS)" -o $(TARGET)/$(BINARY) .

## run — build and run from target/
run: build
	./$(TARGET)/$(BINARY)

## clean — remove target/ and tmp/
clean:
	rm -rf $(TARGET) $(TMP)

## tmp-dir / target-dir — create dirs if missing
tmp-dir:
	@mkdir -p $(TMP)

target-dir:
	@mkdir -p $(TARGET)
