package node

import "github.com/rasmusrygaard/chord/chord"

func OwnsID(n chord.Node, id chord.ID) bool {
	successor := n.Successor()
	if successor.Identifier() < n.Identifier() {
		return id >= n.Identifier() || id < successor.Identifier()
	}

	return id >= n.Identifier() && id < successor.Identifier()
}

type LinkedListNode struct {
	ID         chord.ID
	Prev, Next *LinkedListNode
}

func (n LinkedListNode) Identifier() chord.ID {
	return n.ID
}

func (n LinkedListNode) FindSuccessor(id chord.ID) chord.Node {
	for cur := n.Successor(); ; cur = cur.Successor() {
		if OwnsID(cur, id) {
			return cur
		}

		if cur == n {
			return nil
		}
	}
}

func (n LinkedListNode) Predecessor() chord.Node {
	return n.Prev
}

func (n LinkedListNode) Successor() chord.Node {
	return n.Next
}

func NewLinkedList(id chord.ID) LinkedListNode {
	node := LinkedListNode{
		ID: id,
	}
	node.Prev = &node
	node.Next = &node
	return node
}
