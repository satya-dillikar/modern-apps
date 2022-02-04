package messages

import "strconv"

type JsonRequest struct {
	JsonRequestString string `json:"jsonRequestString"`
}

type JsonResponse struct {
	JsonResponseString string `json:"jsonResponseString"`
}

/* Information/Metadata about node */
type NodeInfo struct {
	NodeId     int    `json:"nodeId"`
	NodeIpAddr string `json:"nodeIpAddr"`
	Port       string `json:"port"`
}

/* Just for pretty printing the node info */
func (node NodeInfo) String() string {
	return "NodeInfo:{ nodeId:" + strconv.Itoa(node.NodeId) +
		", nodeIpAddr:" + node.NodeIpAddr + ", port:" + node.Port + " }"
}
