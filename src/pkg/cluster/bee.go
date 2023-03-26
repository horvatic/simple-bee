package cluster

import (
	"fmt"
	"net/http"

	"github.com/horvatic/simple-bee/pkg/node"
	"github.com/horvatic/simple-bee/pkg/store"
	"github.com/horvatic/simple-bee/pkg/system"
)

var clusterStore store.Store

func Run(newStore store.Store) {
	clusterStore = newStore
	http.HandleFunc("/nodeinfo", nodeInfo)
	http.HandleFunc("/queeninfo", isQueenInfo)
	http.ListenAndServe(":8090", nil)
}

func isQueenInfo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, GetJsonQueenMessage(clusterStore))
}

func nodeInfo(w http.ResponseWriter, req *http.Request) {
	sampleNode := node.NewNode(system.GetLocalIP(), system.GetHostName(), system.GetTags())

	fmt.Fprintf(w, GetJsonNodeMessage(sampleNode))
}
