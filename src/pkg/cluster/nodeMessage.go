package cluster

import (
	"encoding/json"

	"github.com/horvatic/simple-bee/pkg/node"
)

type NodeInfo struct {
	Ip   string   `json:"ip"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func GetJsonNodeMessage(n *node.Node) string {
	nodeMessage, err := json.Marshal(NodeInfo{Ip: n.GetIp(), Name: n.GetName(), Tags: n.GetTags()})
	if err != nil {
		panic(err)
	}

	return string(nodeMessage)
}

func GetJsonNodesMessage(nodes []node.Node) string {
	var nodesInfo []NodeInfo = make([]NodeInfo, 0)
	for _, n := range nodes {
		nodesInfo = append(nodesInfo, NodeInfo{Ip: n.GetIp(), Name: n.GetName(), Tags: n.GetTags()})
	}
	nodeMessage, err := json.Marshal(nodesInfo)
	if err != nil {
		panic(err)
	}

	return string(nodeMessage)
}