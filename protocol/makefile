.PHONY: build

PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go
PROTOC := $(shell which protoc)

ifeq ($(PROTOC),)
    PROTOC = must-rebuild
endif

build: *.pb.go
	@echo -n > spotify/spotify.go
	@echo "package spotify" >> spotify/spotify.go
	@echo >> spotify/spotify.go
	@echo "var Version = \"$(VERSION)\"" >> spotify/spotify.go
	@go fmt spotify/*.go

$(PROTOC):
ifeq ($(UNAME), Darwin)
	@brew install protobuf >>$(LOGGER)
endif
ifeq ($(UNAME), Linux)
	@sudo apt-get -y install protobuf-compiler >>$(LOGGER)
endif

$(PROTOC_GEN_GO):
	@go get -u github.com/golang/protobuf/protoc-gen-go >>$(LOGGER)

*.pb.go: proto/*.proto | $(PROTOC_GEN_GO) $(PROTOC)
ifeq ($(wildcard spotify), )
	@mkdir spotify >>$(LOGGER)
else
	@rm -rf spotify >>$(LOGGER)
	@mkdir spotify >>$(LOGGER)
endif
	@protoc -I proto --go_out=spotify $? >>$(LOGGER)

.PHONY: clean
clean:
ifneq ($(wildcard spotify), )
	@rm -rf spotify >>$(LOGGER)
endif
