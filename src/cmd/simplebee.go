package main

import (
	"github.com/horvatic/simple-bee/pkg/cluster"
	"github.com/horvatic/simple-bee/pkg/store"
)

func main() {
	cluster.Run(store.NewMemoryStore())
}
