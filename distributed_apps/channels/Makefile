
NAME:=dapp-channels
DOCKER_REPOSITORY:=dsatya6
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
VERSION:=0.1.0

.PHONY: build clean test

build :
	go build -o bin/pxcx_main ./*.go

clean :
	@rm bin/*

test:
	bin/pxcx_main

