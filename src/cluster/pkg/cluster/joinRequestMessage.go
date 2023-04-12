package cluster

import (
	"encoding/json"
)

type JoinIdInfo struct {
	JoinId string `json:"joinId"`
}

func requestJoinCluster(id string) string {
	joinMessage, err := json.Marshal(JoinIdInfo{JoinId: id})
	if err != nil {
		panic(err)
	}

	return string(joinMessage)
}
