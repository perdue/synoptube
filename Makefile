# The Makefile compiles all Go programs located in subdirectories of 'cmd'.
# It places the compiled executables in the 'bin' directory.
# The 'all' target rebuilds executables if any .go files change in 'cmd' or 'pkg'.

# Define the output directory for the compiled executables.
GOBIN := bin

# Find all subdirectories inside the 'cmd' directory.
# These subdirectories represent the individual commands to be built.
CMD_DIRS := $(wildcard cmd/*)

# Extract the base names of the command directories to use as executable names.
# For example, 'cmd/server' becomes 'server'.
CMDS := $(notdir $(CMD_DIRS))

# You can define a custom name for your executables here.
# For example, to build cmd/synoptube-cli as 'syt'.
EXECUTABLE_NAME_synoptube-cli := syt

# Get a list of all executable directories.
EXECUTABLE_DIRS := $(notdir $(wildcard cmd/*))

# The list of final executable names to be built.
# This uses a 'patsubst' to substitute the directory name with the custom name.
# If a custom name isn't defined, it falls back to the directory name.
EXECUTABLES := $(foreach dir, $(EXECUTABLE_DIRS), $(or $(EXECUTABLE_NAME_$(dir)), $(dir)))

# Find all Go files in the 'pkg' directory and its subdirectories.
# This ensures that any changes to shared packages trigger a recompile.
CMD_FILES := $(shell find cmd -name "*.go")
PKG_FILES := $(shell find pkg -name "*.go")

# Utility rule to print variable values.
# Example: make print-CMD_FILES.
print-%  : ; @echo $* = $($*)

# Define all targets for the 'all' rule.
# The 'all' rule will build each command and place it in the 'bin' directory.
.PHONY: all
all: build

.PHONY: build
build: $(addprefix $(GOBIN)/, $(CMDS))

# This is a pattern rule to build each individual command.
# The target is the executable in the 'bin' directory.
# The prerequisites are all Go files in the command's own directory,
# as well as all Go files from the 'pkg' directory.
$(GOBIN)/%: $(CMD_FILES) $(PKG_FILES)
	@mkdir -p $(GOBIN)
	go build -o $@ ./cmd/$(notdir $@)

# This target cleans up the compiled executables.
.PHONY: clean
clean:
	rm -rvf $(GOBIN)

# This target runs 'go mod tidy' in the project.
.PHONY: tidy
tidy:
	go mod tidy

# This target runs 'go vet' on all Go files in the project.
.PHONY: vet
vet:
	go vet ./...
