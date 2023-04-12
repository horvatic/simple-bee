package cluster

type ClusterJoinInfo struct {
	Ip   string   `json:"ip"`
	Name string   `json:"Name"`
	Tags []string `json:"Tags"`
	Id string `json:"clusterId"`
}

type JoinClusterPackage struct {
	AuthCode        string `json:"authCode"`
	ClusterQueenUri string `json:"clusterQueenUri"`
}