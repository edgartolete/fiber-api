GOCMD := go
GORUN := $(GOCMD) run
GOBUILD := $(GOCMD) build

run:
	$(GORUN) cmd/app/main.go

build:
	$(GOBUILD) cmd/app/main.go