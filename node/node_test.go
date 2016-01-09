package node_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rasmusrygaard/chord/chord"
	"github.com/rasmusrygaard/chord/node"
)

func TestLinkedListNode(t *testing.T) {
	n := node.NewLinkedList(chord.ID(10))

	assert.True(t, node.OwnsID(n, chord.ID(8)))
	assert.True(t, node.OwnsID(n, chord.ID(10)))
	assert.True(t, node.OwnsID(n, chord.ID(12)))
}
