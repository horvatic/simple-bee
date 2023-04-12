package main

import (
	"os"

	"github.com/horvatic/simple-bee/beekeeper/pkg/commandBuild"
)

func main() {
	com := commandBuild.Build(os.Args[1:])
	com.Run()
}
