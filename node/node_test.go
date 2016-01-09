package node_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rasmusrygaard/chord/chord"
	"github.com/rasmusrygaard/chord/node"
)

func TestLinkedListNode(t *testing.T) {
	first := node.NewLinkedList(chord.ID(10))

	checkMappings(t, map[chord.Node][]chord.ID{
		first: []chord.ID{8, 10, 12},
	})

	second := &node.LinkedListNode{
		ID: chord.ID(20),
	}
	assert.Equal(t, nil, first.Join(second))

	checkMappings(t, map[chord.Node][]chord.ID{
		first:  []chord.ID{1, 40, 10},
		second: []chord.ID{11, 19, 20},
	})

	third := &node.LinkedListNode{
		ID: chord.ID(1),
	}
	assert.Equal(t, nil, first.Join(third))

	checkMappings(t, map[chord.Node][]chord.ID{
		first:  []chord.ID{5, 10},
		second: []chord.ID{11, 19},
		third:  []chord.ID{1, 40},
	})
}

func checkMappings(t *testing.T, mappings map[chord.Node][]chord.ID) {
	for n, ids := range mappings {
		for _, id := range ids {
			assert.True(t, node.OwnsID(n, id), "Expected %d to own %d", n.Identifier(), id)
		}
	}
}
