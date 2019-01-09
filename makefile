.PHONY: all
all: build
	@echo >> /dev/null

ifeq ($(GOPATH),)
GOPATH=$(HOME)/go
endif

ifeq ($(DEBUG),)
BUILDTAGS=release
else
BUILDTAGS=debug
endif

ifeq ($(VERBOSE),)
LOGGER=/dev/null 2>&1
else
LOGGER=/dev/stdout 2>&1
endif

ifeq ($(VERSION),)
VERSION=$(shell git rev-parse --short HEAD)
endif

UNAME=$(shell uname)

export GOPATH
export VERSION
export UNAME
export LOGGER
export BUILDTAGS

define in_progress_msg
	@echo "> $(1)..."
endef

define finish_msg
	@echo "> $(1) finished"
endef

.PHONY: go_mod_verify
go-mod-verify:
	$(call in_progress_msg,"verifying mods")
	@go mod verify >>$(LOGGER)
	$(call finish_msg,"verifying mods")

.PHONY: go_mod_download
go-mod-download:
	$(call in_progress_msg,"downloading mods")
	@go mod download >>$(LOGGER)
	$(call finish_msg,"downloading mods")

.PHONY: go_build
go-build:
	$(call in_progress_msg,"building respot2")
	@go build -v ./... >> $(LOGGER)
	$(call finish_msg,"building respot2")

.PHONY: build_protocol
build_protocol:
	$(call in_progress_msg,"building protocol")
	@cd protocol;./build.sh >>$(LOGGER);cd ..
	$(call finish_msg,"building protocal")

.PHONY: build
build: show_version clean build_protocol go-mod-download go-mod-verify go-build
	$(call finish_msg,"building $(VERSION)")

.PHONY: clean
clean:
	$(call in_progress_msg,"cleaning up protocol")
	@cd protocol;./build.sh clean;cd .. >>$(LOGGER)
	$(call finish_msg,"cleaning up protocol")
	$(call in_progress_msg,"cleaning up build cache")
	@-go clean ./...  >>$(LOGGER)
	$(call finish_msg,"cleaning up build cache")

.PHONY: show_version
show_version:
	$(call in_progress_msg,"building $(VERSION)")