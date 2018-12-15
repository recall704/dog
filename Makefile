
GO           ?= go
BUILD		 ?= $(GO) build
GOOS         ?= linux
GOARCH       ?= arm
GOARM        ?= 7


.PHONY: pi
pi:
	GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) $(BUILD) -v -o bin/dog src/main.go 
