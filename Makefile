# Target binary name
BINARY_NAME := code_scanner

# Source files
SRC := ./cmd/main.go

# Build target
build: 
	go build -o ${BINARY_NAME} ${SRC}

# Clean target
clean:
	rm -f $(BINARY_NAME)
	rm -rf output-dir

# Default target
all: build

.PHONY: build clean all