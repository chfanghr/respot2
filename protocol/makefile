PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go
PROTOC := $(shell which protoc)

ifeq ($(LOGGER),)
LOGGER=/dev/null 2>&1
endif

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

.PHONY: build
build: *.pb.go
	@echo -n > spotify/spotify.go
	@echo "package spotify" >> spotify/spotify.go
	@echo >> spotify/spotify.go
	@echo "var Version = \"$(VERSION)\"" >> spotify/spotify.go
	@go fmt spotify/*.go >> spotify/spotify.go

INCLUDEDIR:=proto/
PROTOFILE:=proto/*.proto
OUTPUTDIR:=./spotify

*.pb.go: $(PROTOFILE)
	@$(PROTOC) -I $(INCLUDEDIR) --plugin=$(PROTOC_GEN_GO) --go_out $(OUTPUTDIR) $(PROTOFILE) >>$(LOGGER)

.PHONY: clean
clean:
	@-rm -rf $(OUTPUTDIR) >>$(LOGGER)

.PHONY: mkdir
mkdir:
	@-mkdir $(OUTPUTDIR) >>$(LOGGER)