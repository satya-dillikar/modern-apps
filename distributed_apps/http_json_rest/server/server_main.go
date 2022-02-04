package main

/* All useful imports */
import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"

	"satya.com/http_json_rest/messages"
)

var me messages.NodeInfo

/* The entry point for our System */
func main() {
	/* Parse the provided parameters on command line */
	clusterip := flag.String("clusterip", "127.0.0.1:9001", "ip address of any node to connnect")
	flag.Parse()

	fmt.Println("Starting http server")
	fmt.Println("cluster ip :", *clusterip)
	startHttpServer()
}

func startHttpServer() {
	fmt.Println("Starting http server.")
	http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/", homeHandler)
	myIpAddrs, _ := net.InterfaceAddrs()
	me = messages.NodeInfo{NodeId: -1, NodeIpAddr: myIpAddrs[0].String(), Port: "9001"}
	http.ListenAndServe(":9001", nil)
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t messages.JsonRequest
	decoder.Decode(&t)
	fmt.Println("Got request string : ", t.JsonRequestString)

	responseJson := messages.JsonResponse{
		JsonResponseString: "query result from server:" + t.JsonRequestString,
	}
	json.NewEncoder(w).Encode(responseJson)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Me : %s", me.String())
}
