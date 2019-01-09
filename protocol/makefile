ifeq ($(LOGGER),)
LOGGER=/dev/stdout 2>&1
endif

ifeq ($(GOPATH),)
LOGGER=$(HOME)/go
endif

PROTOC_GEN_GO= $(GOPATH)/bin/protoc-gen-go
PROTOC= $(shell which protoc)

INCLUDEDIR=proto/
PROTOFILE=proto/*.proto
OUTPUTDIR=./spotify

.PHONY: all
all: deps clean mkdir build
	@echo >> /dev/null

.PHONY: deps
deps: $(PROTOC_GEN_GO) $(PROTOC)

$(PROTOC_GEN_GO):
	@go get -u -v github.com/golang/protobuf/protoc-gen-go >>$(LOGGER)

$(PROTOC):
ifeq ($(UNAME), Darwin)
	@brew install protobuf >>$(LOGGER)
endif
ifeq ($(UNAME), Linux)
	@sudo apt-get -y install protobuf-compiler >>$(LOGGER)
endif
PROTOC= $(shell which protoc)

.PHONY: build
build:
	@$(PROTOC) -I $(INCLUDEDIR) --plugin=$(PROTOC_GEN_GO) --go_out $(OUTPUTDIR) $(PROTOFILE) >>$(LOGGER)
	@echo -n > spotify/spotify.go
	@echo "package spotify" >> spotify/spotify.go
	@echo >> spotify/spotify.go
	@echo "var Version = \"$(VERSION)\"" >> spotify/spotify.go
	@go fmt spotify/*.go >> spotify/spotify.go

.PHONY: clean
clean:
	@-rm -rf $(OUTPUTDIR) >>$(LOGGER)

.PHONY: mkdir
mkdir:
	@-mkdir $(OUTPUTDIR) >>$(LOGGER)