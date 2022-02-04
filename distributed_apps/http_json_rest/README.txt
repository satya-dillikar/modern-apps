Start Http server :
go run server/main.go 

Start Client & type input:
go run  client/main.go --clusterip 127.0.0.1:9001


curl -X POST -d '{"jsonRequestString": "this is a request"}' http://127.0.0.1:9001/query
{"jsonResponseString":"query result from server:this is a request"}
