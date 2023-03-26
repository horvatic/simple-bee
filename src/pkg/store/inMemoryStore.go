package store

type MemoryStore struct {
	isQueen bool
}

func NewMemoryStore() Store {
	s := MemoryStore{isQueen: true}
	return &s
}

func (m MemoryStore) SetIsQueen() {
	m.isQueen = true
}

func (m MemoryStore) SetIsNotQueen() {
	m.isQueen = false
}

func (m MemoryStore) IsQueen() bool {
	return m.isQueen
}
