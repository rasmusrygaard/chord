package chord

type ID int64

type Node interface {
	Identifier() ID
	FindPredecessor(ID) Node
	Successor() Node
}

func (n Node) FindSuccessor(id ID) {
	predecessor := n.FindPredecessor(id)
	return predecessor.Successor()
}

func (n Node) OwnsID(id ID) bool {
	successor := n.Successor()
	if successor.Identifier() < n.Identifier() {
		return id >= n.Identifier() || id < successor.Identifier()
	}

	return id >= n.Identifer() && id < successor.Identifer()
}

type LinkedListNode struct {
	ID         ID
	Prev, Next *LinkedListNode
}

func (n LinkedListNode) Identifier() ID {
	return n.ID
}

func (n LinkedListNode) FindPredecessor(id ID) Node {
	next := n.Successor()
	for cur := n.Successor(); ; cur = cur.Successor() {
		if cur.OwnsID(id) {
			return cur
		}

		if cur == n {
			return nil
		}
	}
}

func (n LinkedListNode) Predecessor() Node {
	return n.Prev
}
