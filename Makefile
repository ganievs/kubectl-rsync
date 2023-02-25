default: build

build:
	go build cmd/kubectl-rsync.go

install:
	go install cmd/kubectl-rsync.go

lint:
	golangci-lint run --fix ./...
