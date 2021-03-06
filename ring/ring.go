package ring

import (
	"github.com/rasmusrygaard/chord/chord"
)

type Ring struct {
	Known chord.Node
}

func (r Ring) Lookup(id chord.ID) (chord.Node, error) {
	return r.Known.FindSuccessor(id), nil
}

func (r Ring) Join(n chord.Node) error {
	return r.Known.Join(n)
}
