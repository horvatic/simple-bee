package joinCommand

import (
	"fmt"
	"io"
	"net/http"
)

type JoinRequest struct {
	ClusterQueenUri string
}

func (j *JoinRequest) Run() error {
	resp, err := http.Get("http://" + j.ClusterQueenUri + "/requestjoincode")
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
