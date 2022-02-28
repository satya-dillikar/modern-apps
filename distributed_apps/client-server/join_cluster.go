package main

/* Al useful imports */
import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

/* Information/Metadata about node */
type NodeInfo struct {
	NodeId     int    `json:"nodeId"`
	NodeIpAddr string `json:"nodeIpAddr"`
	Port       string `json:"port"`
}

/* A standard format for a Request/Response for adding node to cluster */
type AddToClusterMessage struct {
	Source  NodeInfo `json:"source"`
	Dest    NodeInfo `json:"dest"`
	Message string   `json:"message"`
}

/* Just for pretty printing the node info */
func (node NodeInfo) String() string {
	return "NodeInfo:{ nodeId:" + strconv.Itoa(node.NodeId) +
		", nodeIpAddr:" + node.NodeIpAddr + ", port:" + node.Port + " }"
}

/* Just for pretty printing Request/Response info */
func (req AddToClusterMessage) String() string {
	return "AddToClusterMessage:{\n  source:" + req.Source.String() +
		",\n  dest: " + req.Dest.String() + ",\n  message:" + req.Message + " }"
}

/* The entry point for our System */
func main() {
	/* Parse the provided parameters on command line */
	makeMasterOnErrorPtr := flag.Bool("makeMasterOnError", false,
		"make this node master if unable to connect to the cluster ip provided.")
	clusterIpPtr := flag.String("clusterIp", "127.0.0.1:8001", "ip address of any node to connnect")
	myPortPtr := flag.String("myPort", "8001", "ip address to run this node on. default is 8001.")
	flag.Parse()

	/* Generate id for myself */
	rand.Seed(time.Now().UTC().UnixNano())
	myNodeId := rand.Intn(99999999)

	myIpAddrs, _ := net.InterfaceAddrs()
	myNode := NodeInfo{NodeId: myNodeId, NodeIpAddr: myIpAddrs[0].String(), Port: *myPortPtr}
	destNode := NodeInfo{NodeId: -1, NodeIpAddr: strings.Split(*clusterIpPtr, ":")[0],
		Port: strings.Split(*clusterIpPtr, ":")[1]}
	fmt.Println("My details:", myNode.String())

	/* Try to connect to the cluster, and send request to cluster if able to connect */
	ableToConnect := connectToCluster(myNode, destNode)

	/*
	 * Listen for other incoming requests form other nodes to join cluster
	 * Note: We are not doing anything fancy right now to make this node as master. Not yet!
	 */
	if ableToConnect || (!ableToConnect && *makeMasterOnErrorPtr) {
		if *makeMasterOnErrorPtr {
			fmt.Println("Will start this node as master.")
		}
		listenOnPort(myNode)
	} else {
		fmt.Println("Quitting system. Set makeMasterOnErrorPtr flag to make the node master.", myNodeId)
	}
}

/*
 * This is a useful utility to format the json packet to send requests
 * This tiny block is sort of important else you will end up sending blank messages.
 */
func getAddToClusterMessage(source NodeInfo, destNode NodeInfo, message string) AddToClusterMessage {
	return AddToClusterMessage{
		Source: NodeInfo{
			NodeId:     source.NodeId,
			NodeIpAddr: source.NodeIpAddr,
			Port:       source.Port,
		},
		Dest: NodeInfo{
			NodeId:     destNode.NodeId,
			NodeIpAddr: destNode.NodeIpAddr,
			Port:       destNode.Port,
		},
		Message: message,
	}
}

func connectToCluster(myNode NodeInfo, destNode NodeInfo) bool {
	/* connect to this socket details provided */
	connOut, err := net.DialTimeout("tcp", destNode.NodeIpAddr+":"+destNode.Port,
		time.Duration(10)*time.Second)
	if err != nil {
		if _, ok := err.(net.Error); ok {
			fmt.Println("Couldn't connect to cluster.", myNode.NodeId)
			return false
		}
	} else {
		fmt.Println("Connected to cluster. Sending message to node.")
		text := "Hi Buddy.. please add me to the cluster.."
		requestMessage := getAddToClusterMessage(myNode, destNode, text)
		json.NewEncoder(connOut).Encode(&requestMessage)

		decoder := json.NewDecoder(connOut)
		var responseMessage AddToClusterMessage
		decoder.Decode(&responseMessage)
		fmt.Println("\nGot response:\n" + responseMessage.String())

		return true
	}
	return false
}

func listenOnPort(myNode NodeInfo) {
	/* Listen for incoming messages */
	listener, _ := net.Listen("tcp", fmt.Sprint(":"+myNode.Port))
	/* accept connection on port */
	/* not sure if looping infinetely on listener.Accept() is good idea */
	for {
		connIn, err := listener.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				fmt.Println("Error received while listening.", myNode.NodeId)
			}
		} else {
			var requestMessage AddToClusterMessage
			json.NewDecoder(connIn).Decode(&requestMessage)
			fmt.Println("\nGot request:\n" + requestMessage.String())

			text := "Sure Buddy.. too easy.."
			responseMessage := getAddToClusterMessage(myNode, requestMessage.Source, text)
			json.NewEncoder(connIn).Encode(&responseMessage)
			connIn.Close()
		}
	}
}
