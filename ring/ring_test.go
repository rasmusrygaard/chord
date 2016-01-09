package ring_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rasmusrygaard/chord/chord"
	"github.com/rasmusrygaard/chord/node"
	"github.com/rasmusrygaard/chord/ring"
)

func TestRing(t *testing.T) {
	first := node.NewLinkedList(chord.ID(10))
	r := ring.Ring{Known: first}

	for _, id := range []chord.ID{chord.ID(1), chord.ID(10), chord.ID(20)} {
		res, _ := r.Lookup(id)
		assert.Equal(t, first.Identifier(), res.Identifier())
	}

	second := node.LinkedListNode{
		ID: chord.ID(20),
	}
	assert.Nil(t, r.Join(&second))

	first.Dump()

	var res chord.Node
	res, _ = r.Lookup(chord.ID(1))
	assert.Equal(t, first.Identifier(), res.Identifier())

	res, _ = r.Lookup(chord.ID(12))
	assert.Equal(t, second.Identifier(), res.Identifier())

	res, _ = r.Lookup(chord.ID(21))
	assert.Equal(t, first.Identifier(), res.Identifier())
}
