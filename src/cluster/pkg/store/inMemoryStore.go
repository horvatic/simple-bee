package store

type MemoryStore struct {
	isQueen bool
	requestJoinIds []string
}

func NewMemoryStore() Store {
	s := MemoryStore{isQueen: true, requestJoinIds: make([]string, 0)}
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

func (m MemoryStore) SaveJoinRequestId(id string) {
	m.requestJoinIds = append(m.requestJoinIds, id)
}