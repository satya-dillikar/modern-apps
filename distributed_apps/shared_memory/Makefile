
NAME:=dapp-shared-memory
DOCKER_REPOSITORY:=dsatya6
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
VERSION:=0.1.0

.PHONY: build clean test

build :
	go build -o bin/xpxc-mem ./xpxc-mem.go

clean :
	@rm bin/*

test:
	# go test -v ./...
	bin/xpxc-mem