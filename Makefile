GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

BINARY_NAME=mdf

BUILD_DIR=build

LOCAL_GOCACHE=$(PWD)/.gocache

INSTALL_DIR=$(HOME)/.local/bin

SHELL_NAME := $(shell basename $$SHELL)
ifeq ($(SHELL_NAME),bash)
    SHELL_CONFIG := $(HOME)/.bashrc
else ifeq ($(SHELL_NAME),zsh)
    SHELL_CONFIG := $(HOME)/.zshrc
else ifeq ($(SHELL_NAME),fish)
    SHELL_CONFIG := $(HOME)/.config/fish/config.fish
else
    SHELL_CONFIG := $(HOME)/.profile
endif

all: build

build:
	mkdir -p $(BUILD_DIR)
	mkdir -p $(LOCAL_GOCACHE)
	GOCACHE=$(LOCAL_GOCACHE) $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -rf $(LOCAL_GOCACHE)

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

install: build
	mkdir -p $(INSTALL_DIR)
	install -m 755 $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)
	@if [ "$(SHELL_NAME)" = "fish" ]; then \
		echo "set -gx PATH $(INSTALL_DIR) \$$PATH" >> $(SHELL_CONFIG); \
	else \
		echo "export PATH=\"$(INSTALL_DIR):\$$PATH\"" >> $(SHELL_CONFIG); \
	fi
	@echo "PATH has been updated in $(SHELL_CONFIG). Please run 'source $(SHELL_CONFIG)' to apply changes."

uninstall:
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@if [ "$(SHELL_NAME)" = "fish" ]; then \
		sed -i '/set -gx PATH $(subst /,\/,$(INSTALL_DIR)) \$$PATH/d' $(SHELL_CONFIG); \
	else \
		sed -i '/export PATH="$(subst /,\/,$(INSTALL_DIR)):\$$PATH"/d' $(SHELL_CONFIG); \
	fi
	@echo "PATH modification has been removed from $(SHELL_CONFIG). Please run 'source $(SHELL_CONFIG)' to apply changes."

.PHONY: all build clean run install uninstall