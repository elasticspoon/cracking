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

func TestPathsWithSum(t *testing.T) {
	tree := &BinTree{value: 20}

	tree.left = &BinTree{value: 11}
	tree.right = &BinTree{value: 3}

	tree.left.left = &BinTree{value: 14}
	tree.left.right = &BinTree{value: 15}
	tree.right.right = &BinTree{value: 8}

	tree.left.left.left = &BinTree{value: -3}
	tree.left.right.left = &BinTree{value: 9}
	tree.right.right.left = &BinTree{value: 7}
	tree.right.right.right = &BinTree{value: 2}

	tree.right.right.left.right = &BinTree{value: -4}
	tree.left.right.left.right = &BinTree{value: 2}

	tests := []struct {
		target        int
		expectedCount int
	}{
		{target: 20, expectedCount: 1},
		{target: 31, expectedCount: 2},
		{target: 0, expectedCount: 0},
		{target: 11, expectedCount: 5},
		{target: -3, expectedCount: 1},
		{target: 15, expectedCount: 2},
		{target: -83, expectedCount: 0},
	}

	for _, test := range tests {
		count := tree.NumPathsWithSum(test.target)
		if count != test.expectedCount {
			t.Errorf("Expected %v paths with sum %v, but got %v", test.expectedCount, test.target, count)
		}
	}
}
