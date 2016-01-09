package node

import "github.com/rasmusrygaard/chord/chord"

func OwnsID(n chord.Node, id chord.ID) bool {
	if n.Successor().Identifier() == n.Identifier() {
		return true
	}

	pre := n.Predecessor()
	if pre.Identifier() > n.Identifier() {
		return id > pre.Identifier() || id <= n.Identifier()
	}

	return id <= n.Identifier() && id > pre.Identifier()
}
