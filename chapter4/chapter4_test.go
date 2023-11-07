package chapter4

import (
	"testing"
)

func TestRandomNode(t *testing.T) {
	tree := &BinTree{value: 11, numNodes: 10}

	tree.left = &BinTree{value: 7, numNodes: 4}
	tree.left.left = &BinTree{value: 5, numNodes: 1}
	tree.left.right = &BinTree{value: 9, numNodes: 2}
	tree.left.right.left = &BinTree{value: 8, numNodes: 1}

	tree.right = &BinTree{value: 15, numNodes: 5}
	tree.right.right = &BinTree{value: 20, numNodes: 4}

	tree.right.right.left = &BinTree{value: 29, numNodes: 1}
	tree.right.right.right = &BinTree{value: 32, numNodes: 2}

	tree.right.right.right.right = &BinTree{value: 50, numNodes: 1}

	results := make(map[int]int)
	// for i := 0; i < 11000; i++ {
	for i := 0; i < 10000; i++ {
		node := tree.RandomNode()
		results[node.value]++
	}

	for k, v := range results {
		if v < 900 || v > 1100 {
			t.Errorf("Expected %v to be selected at least 1000 times, but was selected %v times", k, v)
		}
	}
}
