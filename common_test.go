package container

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	nodes := []*Node{
		{100, "a"},
		{30, "b"},
		{50, "c"},
		{20, "d"},
	}

	q := NewPriorityQueue(nodes)
	q.Push(&Node{200, "f"})
	q.Push(&Node{2, "g"})
	for q.Len() > 0 {
		node := q.Pop()
		t.Logf("value -> %s, prior -> %d\n", node.Value, node.Prior)
	}
}
