package commandBuild

import (
	"github.com/horvatic/simple-bee/beekeeper/pkg/joinCommand"
)

func Build(command []string) CommandPackage {
	if len(command) == 0 {
		panic("no args passed")
	}

	if command[0] == "-r" {
		if len(command) < 2 {
			panic("Missing auth code")
		}
		return &joinCommand.JoinRequest{ClusterQueenUri: command[1]}
	}

	if command[0] == "-j" {
		if len(command) < 4 {
			panic("Missing auth code")
		}
		return &joinCommand.JoinCluster{AuthCode: command[1], ClusterWorkerUri: command[2], ClusterQueenUri: command[3]}
	}

	panic("Can could process arg")
}
