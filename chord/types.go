package chord

type ID int64

type Node interface {
	Identifier() ID
	FindSuccessor(ID) Node
	Successor() Node
}
