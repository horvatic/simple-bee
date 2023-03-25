package main

import (
	"fmt"

	"github.com/horvatic/simple-bee/pkg/node"
)

func main() {
	sampleTags := []string{"arm", "x64"}
	sampleNode := node.NewNode("1.1.1.0", "test", sampleTags)

	fmt.Println(sampleNode.GetIp())
	fmt.Println(sampleNode.GetName())
	fmt.Println(sampleNode.GetTags())
}
