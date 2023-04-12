package main

import (
	"os"

	"github.com/horvatic/simple-bee/pkg/cluster"
	"github.com/horvatic/simple-bee/pkg/store"
)

func main() {
	port := os.Getenv("CLUSTER_PORT")
	cluster.Run(port, store.NewMemoryStore())
}
