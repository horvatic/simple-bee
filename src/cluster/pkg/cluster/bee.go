package cluster

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/horvatic/simple-bee/pkg/node"
	"github.com/horvatic/simple-bee/pkg/store"
	"github.com/horvatic/simple-bee/pkg/system"
)

var clusterStore store.Store

func Run(port string, newStore store.Store) {
	clusterStore = newStore
	http.HandleFunc("/nodeinfo", nodeInfo)
	http.HandleFunc("/queeninfo", isQueenInfo)
	http.HandleFunc("/clusterjoin", clusterJoin)
	http.HandleFunc("/requestjoincode", requestJoinCode)

	fmt.Println("LISTENING ON PORT: " + port)

	http.ListenAndServe(":" + port, nil)
}

func clusterJoin(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
    var c JoinClusterPackage
    err := decoder.Decode(&c)
    if err != nil {
        panic(err)
    }
	fmt.Println(c)
}

func requestJoinCode(w http.ResponseWriter, req *http.Request) {
	joinId := uuid.NewString()
	fmt.Fprintf(w, requestJoinCluster(joinId))
}

func isQueenInfo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, GetJsonQueenMessage(clusterStore))
}

func nodeInfo(w http.ResponseWriter, req *http.Request) {
	sampleNode := node.NewNode(system.GetLocalIP(), system.GetHostName(), system.GetTags())

	fmt.Fprintf(w, GetJsonNodeMessage(sampleNode))
}
