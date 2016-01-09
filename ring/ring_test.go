package ring_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rasmusrygaard/chord/chord"
	"github.com/rasmusrygaard/chord/node"
	"github.com/rasmusrygaard/chord/ring"
)

func TestRing(t *testing.T) {
	n := node.NewLinkedList(chord.ID(10))
	r := ring.Ring{Known: n}

	for _, id := range []chord.ID{chord.ID(1), chord.ID(10), chord.ID(20)} {
		res, _ := r.Lookup(id)
		assert.Equal(t, n.Identifier(), res.Identifier())
	}
}
