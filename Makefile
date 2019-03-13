# Go related variables.
GOBASE=$(shell pwd)
GOSRC=$(GOBASE)/src
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard $(GOSRC)/*.go)
GONAME=$(shell basename "$(PWD)")

build:
	mkdir $(GOBASE)/bin
	@echo "Building $(GOFILES) to $(GOBASE)/bin"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GOBIN)/$(GONAME) $(GOFILES)

get:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(GOSRC)

clear:
	@clear

clean:
	@echo "Cleaning"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean
	rm -rf $(GOBIN)/$(GONAME)
