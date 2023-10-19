package builtins

import (
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	tests := []struct {
		input     []int
		traversal string
		expected  []int
	}{
		{[]int{5, 3, 7, 2, 4, 6, 8}, "inorder", []int{2, 3, 4, 5, 6, 7, 8}},
		{[]int{5, 3, 7, 2, 4, 6, 8}, "inorder-rec", []int{2, 3, 4, 5, 6, 7, 8}},
		{[]int{5, 3, 7, 2, 4, 6, 8}, "preorder", []int{5, 3, 2, 4, 7, 6, 8}},
		{[]int{5, 3, 7, 2, 4, 6, 8}, "preorder-rec", []int{5, 3, 2, 4, 7, 6, 8}},
		{[]int{5, 3, 7, 2, 4, 6, 8}, "postorder", []int{2, 4, 3, 6, 8, 7, 5}},
		{[]int{5, 3, 7, 2, 4, 6, 8}, "postorder-rec", []int{2, 4, 3, 6, 8, 7, 5}},
	}

	for _, test := range tests {
		tree := NewBinarySearchTree(test.input)
		var result []int

		switch test.traversal {
		case "inorder":
			result = tree.Inorder()
		case "inorder-rec":
			result = tree.InorderRec()
		case "preorder":
			result = tree.PreOrder()
		case "preorder-rec":
			result = tree.PreOrderRec()
		case "postorder":
			result = tree.PostOrder()
		case "postorder-rec":
			result = tree.PostOrderRec()
		}

		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("Expected %s traversal of %v to be %v, got %v", test.traversal, test.input, test.expected, result)
			}
		}
	}
}

func TestBSTDelete(t *testing.T) {
	tests := []struct {
		input    []int
		delete   int
		expected []int
	}{
		// {[]int{5, 3, 7, 2, 4, 6, 8}, 8, []int{2, 3, 4, 5, 6, 7, 8}},
		{[]int{5, 3, 7, 2, 4, 6, 8}, 8, []int{2, 3, 4, 5, 6, 7}},
		{[]int{5, 3, 7, 2, 4, 6, 8}, 6, []int{2, 3, 4, 5, 7, 8}},
		{[]int{5, 3, 7, 2, 4, 6}, 7, []int{2, 3, 4, 5, 6}},
		{[]int{5, 3, 7, 2, 4, 8}, 7, []int{2, 3, 4, 5, 8}},
		{[]int{5, 3, 7, 2, 4, 6, 8}, 7, []int{2, 3, 4, 5, 6, 8}}, // Delete node with 2 children, right child has no left child
		{[]int{5, 3, 7, 2, 4, 6, 8}, 3, []int{2, 4, 5, 6, 7, 8}},
		{[]int{5, 3, 7, 2, 4, 6, 8}, 5, []int{2, 3, 4, 6, 7, 8}},
	}

	for _, test := range tests {
		tree := NewBinarySearchTree(test.input)
		tree.Delete(test.delete)
		result := tree.Inorder()

		if len(result) != len(test.expected) {
			t.Fatalf("Expected %v to be %v when deleting %d", result, test.expected, test.delete)
		}

		for i, v := range result {
			if test.expected[i] != v {
				t.Errorf("Expected %v to be %v when deleting %d", result, test.expected, test.delete)
				break
			}
		}
	}
}
