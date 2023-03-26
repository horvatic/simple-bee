package cluster

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/horvatic/simple-bee/pkg/node"
	"github.com/horvatic/simple-bee/pkg/store"
	"github.com/horvatic/simple-bee/pkg/system"
)

var clusterStore store.Store

func Run(newStore store.Store) {
	clusterStore = newStore
	http.HandleFunc("/nodeinfo", nodeInfo)
	http.HandleFunc("/queeninfo", isQueenInfo)
	http.HandleFunc("/addworker", addWorkerNode)
	http.HandleFunc("/getworkers", getWorkerNodes)
	http.ListenAndServe(":8090", nil)
}

func getWorkerNodes(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, GetJsonNodesMessage(clusterStore.GetAllWorkerNodes()))
}

func addWorkerNode(w http.ResponseWriter, req *http.Request) {
	var n NodeInfo
	err := json.NewDecoder(req.Body).Decode(&n)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	clusterStore.AddWorkNode(node.NewNode(n.Ip, n.Name, n.Tags))
}

func isQueenInfo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, GetJsonQueenMessage(clusterStore))
}

func nodeInfo(w http.ResponseWriter, req *http.Request) {
	sampleNode := node.NewNode(system.GetLocalIP(), system.GetHostName(), system.GetTags())

	fmt.Fprintf(w, GetJsonNodeMessage(sampleNode))
}
