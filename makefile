GOCMD := go
GORUN := $(GOCMD) run
GOBUILD := $(GOCMD) build

run:
	$(GORUN) cmd/api/main.go

build:
	$(GOBUILD) cmd/api/main.go