
NAME:=dapp-client-server
DOCKER_REPOSITORY:=dsatya6
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
VERSION:=0.1.0

.PHONY: build clean test

build :
	go build -o bin/join_cluster ./join_cluster.go

clean :
	@rm bin/*

test:
	# Node 1
	bin/join_cluster --makeMasterOnError &
	@sleep 2
	# Node 2
	bin/join_cluster --myPort 8002 --clusterIp 127.0.0.1:8001 &
	@sleep 2
	# Node 3
	bin/join_cluster --myPort 8004 --clusterIp 127.0.0.1:8001 &
	@sleep 2
	# Node 4
	bin/join_cluster --myPort 8003 --clusterIp 127.0.0.1:8002 &
	@sleep 2
