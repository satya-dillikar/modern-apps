
NAME:=http_json_rest
DOCKER_REPOSITORY:=dsatya6
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
VERSION:=0.1.0

.PHONY: build clean test build-container push-container test-container

build :
	go build -o bin/server_main ./server/server_main.go
	go build -o bin/client_main ./client/client_main.go
clean :
	@rm bin/*

test:
	go test -v ./...  

build-container :
	@docker build -t dsatya6/http_json_rest:$(VERSION) .

push-container :
	@docker push dsatya6/http_json_rest:$(VERSION)	

test-container :
	@docker rm -f http_json_rest || true
	@docker run -dp 9001:9001 --name=http_json_rest $(DOCKER_IMAGE_NAME):$(VERSION)
	@docker ps
	curl -s localhost:9001 | grep Me	