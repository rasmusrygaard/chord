package node

import (
	"fmt"

	"github.com/rasmusrygaard/chord/chord"
)

type LinkedListNode struct {
	ID         chord.ID
	Next, Prev *LinkedListNode
}

func (n LinkedListNode) Identifier() chord.ID {
	return n.ID
}

func (n LinkedListNode) FindSuccessor(id chord.ID) chord.Node {
	fmt.Printf("Looking for %d", id)
	for cur := n.Successor(); ; cur = cur.Successor() {
		fmt.Printf("cur: %d\n", cur.Identifier())
		if OwnsID(cur, id) {
			return cur
		}

		if cur.Identifier() == n.Identifier() {
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

func (n *LinkedListNode) Join(new chord.Node) error {
	newAsLinked, ok := new.(*LinkedListNode)
	if !ok {
		return fmt.Errorf("Attempted to join with a non linked-list node")
	}

	succ := n.FindSuccessor(new.Identifier())
	asLinked := succ.(*LinkedListNode)
	asLinked.insertBefore(newAsLinked)

	return nil
}

func (n *LinkedListNode) insertBefore(pred *LinkedListNode) {
	pred.Next = n
	pred.Prev = n.Prev

	n.Prev.Next = pred
	n.Prev = pred
}

func NewLinkedList(id chord.ID) *LinkedListNode {
	node := LinkedListNode{
		ID: id,
	}
	node.Prev = &node
	node.Next = &node
	return &node
}

func (n *LinkedListNode) Dump() {
	fmt.Printf("%d", n.Identifier())
	for cur := n.Successor(); cur != n; cur = cur.Successor() {
		fmt.Printf(" %d", cur.Identifier())
	}
	fmt.Printf("\n")
}
