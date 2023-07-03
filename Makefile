SHELL = /bin/bash

PWD ?= .
BUILDDIR ?= $(PWD)/app/contract


.PHONY: build
build: fmt vet
	go build -o ./build/vmpx .


.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...


.PHONY: abi
abi:
	@rm -f ./app/contract/*.go;
	@echo "abigen...";
	for i in vmpx; do \
		abigen --abi=$(BUILDDIR)/$$i.abi --type=$$i --pkg=contract > $(BUILDDIR)/$$i.go; \
	done

