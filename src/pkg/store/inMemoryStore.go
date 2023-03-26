package store

import (
	"github.com/horvatic/simple-bee/pkg/node"
)

type MemoryStore struct {
	isQueen bool
	nodes []node.Node
}

func NewMemoryStore() Store {
	s := MemoryStore{isQueen: true, nodes: make([]node.Node, 0)}
	return &s
}

func (m *MemoryStore) SetIsQueen() {
	m.isQueen = true
}

func (m *MemoryStore) SetIsNotQueen() {
	m.isQueen = false
}

func (m *MemoryStore) IsQueen() bool {
	return m.isQueen
}

func (m *MemoryStore) AddWorkNode(n *node.Node) {	
	m.nodes = append(m.nodes, *n)
}

func (m *MemoryStore) GetAllWorkerNodes() []node.Node {
	return m.nodes
}