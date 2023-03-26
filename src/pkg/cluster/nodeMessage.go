package cluster

import (
	"encoding/json"

	"github.com/horvatic/simple-bee/pkg/node"
)

type NodeInfo struct {
	Ip   string   `json:"ip"`
	Name string   `json:"Name"`
	Tags []string `json:"Tags"`
}

func GetJsonNodeMessage(n *node.Node) string {
	nodeMessage, err := json.Marshal(NodeInfo{Ip: n.GetIp(), Name: n.GetName(), Tags: n.GetTags()})
	if err != nil {
		panic(err)
	}

	return string(nodeMessage)
}
