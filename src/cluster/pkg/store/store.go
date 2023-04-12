package store

type Store interface {
	SetIsQueen()
	SetIsNotQueen()
	IsQueen() bool
	SaveJoinRequestId(id string)
}
