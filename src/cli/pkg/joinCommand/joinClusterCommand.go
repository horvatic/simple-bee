package joinCommand

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JoinCluster struct {
	AuthCode         string
	ClusterWorkerUri string
	ClusterQueenUri  string
}

type joinClusterPackage struct {
	AuthCode        string `json:"authCode"`
	ClusterQueenUri string `json:"clusterQueenUri"`
}

func (j *JoinCluster) Run() error {
	joinPackage, _ := json.Marshal(joinClusterPackage{AuthCode: j.AuthCode, ClusterQueenUri: j.ClusterQueenUri})
	fmt.Println(string(joinPackage))

	resp, err := http.Post("http://"+j.ClusterWorkerUri+"/clusterjoin", "application/json", bytes.NewBuffer(joinPackage))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
