package cluster

import (
	"encoding/json"

	"github.com/horvatic/simple-bee/pkg/store"
)

type QueenInfo struct {
	IsQueen bool `json:"isQueen"`
}

func GetJsonQueenMessage(clusterStore store.Store) string {
	queenMessage, err := json.Marshal(QueenInfo{IsQueen: clusterStore.IsQueen()})
	if err != nil {
		panic(err)
	}

	return string(queenMessage)
}
