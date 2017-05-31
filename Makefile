NAME := go-mailmag-sample

## Setup
setup:
	go get github.com/Masterminds/glide
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: setup help
