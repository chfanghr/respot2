BUILD=$(shell git rev-parse --short HEAD)
BUILD_TIME=`date "+%Y%m%d%H%M%S"`
MODULE_NAME=github.com/chfanghr/respot2
PROTOC_GEN_GO=$(GOPATH)/bin/protoc-gen-go
PROTOC=$(shell which protoc)
UNAME=$(shell uname)

ifneq ($(VERSION),)
	VERSION=UNKNOWN
endif

ifeq ($(BUILD_MODE),)
	BUILD_MODE=RELEASE
else 
	BUILD_MODE=DEBUG
endif

ifeq ($(GOROOT),)
	GOROOT=/usr/local/go
endif	
	
ifeq ($(GOPATH),)
	GOPATH=$(HOME)/go
endif

GO=$(GOROOT)/bin/go

ifeq ($(VERBOSE),)
	COMMAND=command 2>&- 1>&-  
else
	COMMAND=set -x&&command
endif

ifeq ($(PROTOC),)
    PROTOC=/usr/bin/protoc
endif

.PHONY:all
all:install test

.PHONY:install
install:build
	@echo "installing packages..."
	@$(COMMAND) $(GO) install $(MODULE_NAME)

.PHONY:clean
clean:clean_build_cache clean_proto
	@echo "cleaning version.go..."
	@-$(COMMAND) rm -rf version.go	
	
.PHONY:build
build:
	@echo "building..."
build:clean_build_cache gen_proto version.go download_mod verify_mod build_all

.PHONY:fmt
fmt:
	@$(GO) fmt $(MODULE_NAME)...

.PHONY:test
test:
	@echo "testing packages..."
	@$(COMMAND) go test $(MODULE_NAME)/...
	
.PHONY:show_dep 
show_dep_mod:
	@$(GO) mod graph
	
.PHONY:download_mod
download_mod:
	@echo "downloading modules..."
	@$(COMMAND) $(GO) mod download

.PHONY:verify_mod
verify_mod:
	@echo "verifying modules..."
	@$(COMMAND) $(GO) mod verify
	 
.PHONY:build_all
build_all:
	@echo "building all packages..."
	@$(COMMAND) $(GO) build -tags $(BUILD_MODE) -v $(MODULE_NAME)/... 
	
.PHONY:clean_build_cache
clean_build_cache:
	@echo "cleaning build cache..."
	@-$(COMMAND) $(GO) clean -i -r -cache -testcache $(MODULE_NAME)/...

.PHONY:gen_proto
gen_proto: 
	@command echo "generating proto files..." 
gen_proto:clean_proto protocol/spotify/*.pb.go

.PHONY:clean_proto
clean_proto:
	@command echo "cleaning proto files..."
	@-$(COMMAND) rm -rf protocol/spotify
	
protocol/spotify/*.pb.go:protocol/proto/*.proto | $(PROTOC) $(PROTOC_GEN_GO)
	@-$(COMMAND) mkdir protocol/spotify
	@command echo "waiting for protoc..."
ifeq ($(VERBOSE),)
	@command $(PROTOC)  -I protocol/proto/ --plugin=$(PROTOC_GEN_GO) --go_out protocol/spotify protocol/proto/*.proto >/dev/null 2>&1
else
	@set -x&&command $(PROTOC)  -I protocol/proto/ --plugin=$(PROTOC_GEN_GO) --go_out protocol/spotify protocol/proto/*.proto
endif

$(PROTOC):
ifeq ($(UNAME), Darwin)
	@$(COMMAND) brew install protobuf
endif
ifeq ($(UNAME), Linux)
	@$(COMMAND) sudo apt -y install protobuf-compiler
endif

$(PROTOC_GEN_GO):
	@$(COMMAND) $(GO) get -u github.com/golang/protobuf/protoc-gen-go
	
version.go:
	@echo -n > version.go
	@printf "package respot2\n\n" >> version.go 
	@printf "const VERSION string = \"%s\"\n" $(VERSION) >> version.go
	@printf "const Build string = \"%s\"\n" $(BUILD) >> version.go
	@printf "const BuildTime string = \"%s\"\n" $(BUILD_TIME) >> version.go
	@$(COMMAND) $(GO) fmt version.go
		
