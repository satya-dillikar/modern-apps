
NAME:=dapp-http-json-rest
DOCKER_REPOSITORY:=dsatya6
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
VERSION:=0.1.0

.PHONY: build clean test build-container push-container test-container

build :
	go build -o bin/server_main ./server/server_main.go
	go build -o bin/client_main ./client/client_main.go

clean :
	@rm bin/*
	@rm -rf Godeps
	@rm -rf go.sum
	@go mod tidy

test:
	go test -v ./...
	./bin/server_main &
	@sleep 5
	# ./bin/client_main
	# @sleep 5
	curl -s localhost:9001 | grep Me

build-container :
	@docker build -t dsatya6/$(NAME):$(VERSION) .

push-container :
	@docker push dsatya6/$(NAME):$(VERSION)

test-container :
	@docker rm -f $(NAME) || true
	@docker run -dp 9001:9001 --name=$(NAME) $(DOCKER_IMAGE_NAME):$(VERSION)
	@docker ps
	@sleep 5
	@echo  ""
	curl -s localhost:9001 | grep Me
	@echo  ""
	@sleep 2
	@docker stop $(NAME) || true
	@docker rm -f $(NAME) || true	