package main

import (
	"fmt"

	"github.com/horvatic/simple-bee/pkg/node"
	"github.com/horvatic/simple-bee/pkg/system"
)

func main() {
	sampleNode := node.NewNode(system.GetLocalIP(), system.GetHostName(), system.GetTags())

	fmt.Println(sampleNode.GetIp())
	fmt.Println(sampleNode.GetName())
	fmt.Println(sampleNode.GetTags())
}
