.PHONY: build_protocol

VERSION:=$(shell git rev-parse --default HEAD)
PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go
PROTOC := $(shell which protoc)
UNAME := $(shell uname)

ifeq ($(PROTOC),)
    PROTOC = must-rebuild
endif

build_protocol: *.pb.go
	@echo -n > spotify/spotify.go
	@echo "package spotify" >> spotify/spotify.go
	@echo >> spotify/spotify.go
	@echo "var Version = \"$(VERSION)\"" >> spotify/spotify.go
	@go fmt spotify/*.go

$(PROTOC):
ifeq ($(UNAME), Darwin)
	@brew install protobuf
endif
ifeq ($(UNAME), Linux)
	@sudo apt-get -y install protobuf-compiler
endif

$(PROTOC_GEN_GO):
	@go get -u github.com/golang/protobuf/protoc-gen-go

*.pb.go: proto/*.proto | $(PROTOC_GEN_GO) $(PROTOC)
ifeq ($(wildcard spotify), )
	@mkdir spotify
else
	@rm -rf spotify
	@mkdir spotify
endif
	@protoc -I proto --go_out=spotify $? >/dev/null 2>&1

.PHONY: clean
clean:
ifneq ($(wildcard spotify), )
	@rm -rf spotify
endif
