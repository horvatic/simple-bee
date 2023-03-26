package store

import (
	"github.com/horvatic/simple-bee/pkg/node"
)

type Store interface {
	SetIsQueen()
	SetIsNotQueen()
	IsQueen() bool
	AddWorkNode(*node.Node)
	GetAllWorkerNodes() []node.Node
}
