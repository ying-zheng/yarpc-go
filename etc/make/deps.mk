# List of vendored go executables needed for make generate
GEN_GO_BIN_DEPS = \
	github.com/golang/mock/mockgen \
	github.com/golang/protobuf/protoc-gen-go \
	github.com/gogo/protobuf/protoc-gen-gogoslick \
	github.com/uber/tchannel-go/thrift/thrift-gen \
	golang.org/x/tools/cmd/stringer \
	go.uber.org/thriftrw \
	go.uber.org/tools/update-license

# List of vendored go executables needed for linting. These are not installed
# automatically and must be requested by $(BIN)/$(basename importPath).

THRIFT_VERSION := 1.0.0-dev
PROTOC_VERSION := 3.15.0
PROGO_GRPC_VERSION := 1.2.0
RAGEL_VERSION := 6.10
ERRCHECK_VERSION := 1.7.0
GOLINT_VERSION := 0.0.0-20210508222113-6edffad5e616
STATICHCHECK_VERSION := 0.5.1
GOIMPORTS_VERSION := 0.24.0

THRIFT_OS := $(UNAME_OS)
PROTOC_OS := $(UNAME_OS)
RAGEL_OS := $(UNAME_OS)

THRIFT_ARCH := $(UNAME_ARCH)
PROTOC_ARCH := $(UNAME_ARCH)
RAGEL_ARCH := $(UNAME_ARCH)

# Use the x86_64 version on arm64 macs.
ifeq ($(UNAME_OS),Darwin)
ifeq ($(UNAME_ARCH),arm64)
THRIFT_ARCH := x86_64
RAGEL_ARCH := x86_64
PROTOC_ARCH := x86_64
endif
endif

# Protoc: use "osx" link for macos
ifeq ($(UNAME_OS),Darwin)
PROTOC_OS := osx
endif

THRIFT_LIB = $(LIB)/thrift-$(THRIFT_VERSION)
THRIFT_TAR = $(THRIFT_LIB)/thrift.tar.gz
THRIFT = $(BIN)/thrift
PROTOC_LIB = $(LIB)/protoc-$(PROTOC_VERSION)
PROTOC_ZIP = $(PROTOC_LIB)/protoc.zip
PROTOC = $(BIN)/protoc
PROTOC_GRPC = $(BIN)/protoc-gen-go-grpc
RAGEL_LIB = $(LIB)/ragel-$(RAGEL_VERSION)
RAGEL_BIN = $(RAGEL_LIB)/ragel
RAGEL = $(BIN)/ragel

GEN_BINS = $(THRIFT) $(PROTOC) $(PROTOC_GRPC) $(RAGEL)

$(RAGEL_BIN):
	@mkdir -p $(RAGEL_LIB)
	curl -L "https://github.com/yarpc/ragel/releases/download/v$(RAGEL_VERSION)/ragel-$(RAGEL_OS)-$(RAGEL_ARCH)" > $(RAGEL_BIN)

$(RAGEL): $(RAGEL_BIN)
	@mkdir -p $(BIN)
	cp $(RAGEL_BIN) $(RAGEL)
	@chmod +x $(RAGEL)

$(THRIFT_TAR):
	@mkdir -p $(THRIFT_LIB)
	curl -L "https://github.com/uber/tchannel-go/releases/download/thrift-v$(THRIFT_VERSION)/thrift-1-$(THRIFT_OS)-$(THRIFT_ARCH).tar.gz" > $(THRIFT_TAR)

$(THRIFT): $(THRIFT_TAR)
	@mkdir -p $(BIN)
	cd $(THRIFT_LIB); tar xzf $(THRIFT_TAR)
	cp $(THRIFT_LIB)/thrift-1 $(THRIFT)

$(PROTOC_ZIP):
	@mkdir -p $(PROTOC_LIB)
	curl -L "https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(PROTOC_OS)-$(PROTOC_ARCH).zip" > $(PROTOC_ZIP)

$(PROTOC): $(PROTOC_ZIP)
	@mkdir -p $(BIN)
	cd $(PROTOC_LIB); unzip $(PROTOC_ZIP)
	cp $(PROTOC_LIB)/bin/protoc $(PROTOC)

$(PROTOC_GRPC):
	@mkdir -p $(BIN)
	GOBIN=$(BIN) go install "google.golang.org/grpc/cmd/protoc-gen-go-grpc@v$(PROGO_GRPC_VERSION)"

$(BIN)/errcheck:
	@mkdir -p $(BIN)
	GOBIN=$(BIN) go install "github.com/kisielk/errcheck@v$(ERRCHECK_VERSION)"

$(BIN)/golint:
	@mkdir -p $(BIN)
	GOBIN=$(BIN) go install "golang.org/x/lint/golint@v$(GOLINT_VERSION)"

$(BIN)/staticcheck:
	@mkdir -p $(BIN)
	GOBIN=$(BIN) go install "honnef.co/go/tools/cmd/staticcheck@v$(STATICHCHECK_VERSION)"

$(BIN)/goimports:
	@mkdir -p $(BIN)
	GOBIN=$(BIN) go install "golang.org/x/tools/cmd/goimports@v$(GOIMPORTS_VERSION)"

define generatedeprule
GEN_BINS += $(BIN)/$(shell basename $1)
endef

define deprule
ifdef SUPPRESS_DOCKER
$(BIN)/$(shell basename $1): go.mod
	@mkdir -p $(BIN)
	PATH=$(BIN):$(PATH) ./etc/bin/vendor-build.sh $(BIN) $1
else
$(BIN)/$(shell basename $1): go.mod
	@mkdir -p $(BIN)
	PATH=$(BIN):$(PATH) ./etc/bin/vendor-build.sh $(BIN) $1
endif
endef

$(foreach i,$(GEN_GO_BIN_DEPS),$(eval $(call generatedeprule,$(i))))
$(foreach i,$(GEN_GO_BIN_DEPS),$(eval $(call deprule,$(i))))

THRIFTRW = $(BIN)/thriftrw
GOLINT = $(BIN)/golint
ERRCHECK = $(BIN)/errcheck
STATICCHECK = $(BIN)/staticcheck
GOIMPORTS = $(BIN)/goimports

.PHONY: predeps
predeps: $(THRIFT) $(PROTOC) $(RAGEL)

.PHONY: deps
deps: predeps $(GEN_BINS) ## install all dependencies
