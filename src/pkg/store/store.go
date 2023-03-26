package store

type Store interface {
	SetIsQueen()
	SetIsNotQueen()
	IsQueen() bool
}
