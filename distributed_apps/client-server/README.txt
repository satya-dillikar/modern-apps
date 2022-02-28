go mod init satya.com/join_cluster
go mod tidy
go build .

Node 1 
$ ./join_cluster --makeMasterOnError

Node 2
$ ./join_cluster --myPort 8002 --clusterIp 127.0.0.1:8001

Node 3
$ ./join_cluster --myPort 8004 --clusterIp 127.0.0.1:8001

Node 4
$ ./join_cluster --myPort 8003 --clusterIp 127.0.0.1:8002
