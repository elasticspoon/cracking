package builtins

type TreeNode struct {
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
	value  int
}

func (n *TreeNode) nodeMin() *TreeNode {
	for n.left != nil {
		n = n.left
	}

	return n
}

func (n *TreeNode) insert(value int) {
	if value < n.value {
		if n.left == nil {
			n.left = &TreeNode{value: value, parent: n}
		} else {
			n.left.insert(value)
		}
	} else if value > n.value {
		if n.right == nil {
			n.right = &TreeNode{value: value, parent: n}
		} else {
			n.right.insert(value)
		}
	}
}

func (n *TreeNode) contains(value int) bool {
	if n.value == value {
		return true
	} else if value < n.value {
		if n.left == nil {
			return false
		} else {
			return n.left.contains(value)
		}
	} else {
		if n.right == nil {
			return false
		} else {
			return n.right.contains(value)
		}
	}
}

func (n *TreeNode) inOrder(res []int) []int {
	if n.left != nil {
		res = n.left.inOrder(res)
	}
	res = append(res, n.value)
	if n.right != nil {
		res = n.right.inOrder(res)
	}
	return res
}

func (n *TreeNode) preOrder(res []int) []int {
	res = append(res, n.value)
	if n.left != nil {
		res = n.left.preOrder(res)
	}
	if n.right != nil {
		res = n.right.preOrder(res)
	}
	return res
}

func (n *TreeNode) postOrder(res []int) []int {
	if n.left != nil {
		res = n.left.postOrder(res)
	}
	if n.right != nil {
		res = n.right.postOrder(res)
	}
	res = append(res, n.value)
	return res
}

type BinarySearchTree struct {
	root *TreeNode
}

func (t *BinarySearchTree) transplant(dest, target *TreeNode) {
	switch dest {
	case t.root:
		t.root = target
	case dest.parent.left:
		dest.parent.left = target
	case dest.parent.right:
		dest.parent.right = target
	default:
		panic("I dont know how we got in this state")
	}

	if target != nil {
		target.parent = dest.parent
	}
}

func (t *BinarySearchTree) deleteNode(target *TreeNode) {
	// case 1: no left child
	if target.left == nil {
		t.transplant(target, target.right)
	} else if target.right == nil {
		t.transplant(target, target.left)
	} else {
		minNode := target.right.nodeMin()
		// if the right child has no left child it is the min node.
		// in that case we simply replace the current node with right child.
		// no changes need to be made to the right subtree of the child
		if target.right != minNode {
			// replace minNode with its right child. This takes sets corrent parentage for right child
			t.transplant(minNode, minNode.right)
			// fix the right nodes of the tree to get transplanted
			minNode.right = target.right
			minNode.right.parent = minNode
		}
		t.transplant(target, minNode)
		// fix the left nodes of the tree to get transplanted
		minNode.left = target.left
		minNode.left.parent = minNode

	}
}

func (t *BinarySearchTree) Delete(value int) {
	tar := t.findNode(value)
	t.deleteNode(tar)
}

func (t *BinarySearchTree) Insert(value int) {
	if t.root == nil {
		t.root = &TreeNode{value: value}
		return
	}
	t.root.insert(value)
}

func (t *BinarySearchTree) Contains(value int) bool {
	if t.root == nil {
		return false
	}
	return t.root.contains(value)
}

func (t *BinarySearchTree) findNode(value int) *TreeNode {
	if t.root == nil {
		return nil
	}
	curr := t.root
	for curr != nil {
		if curr.value == value {
			return curr
		} else if value < curr.value {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return nil
}

func NewBinarySearchTree(vals []int) *BinarySearchTree {
	tree := &BinarySearchTree{}

	for _, v := range vals {
		tree.Insert(v)
	}

	return tree
}

func (t *BinarySearchTree) PostOrderRec() []int {
	res := []int{}

	return t.root.postOrder(res)
}

func (t *BinarySearchTree) PreOrderRec() []int {
	res := []int{}

	return t.root.preOrder(res)
}

func (t *BinarySearchTree) InorderRec() []int {
	res := []int{}

	return t.root.inOrder(res)
}

func (t *BinarySearchTree) PostOrder() []int {
	res := []int{}
	stack := NewStack()

	stack.Push(t.root)
	for stack.Len() > 0 {
		curr := stack.Pop()
		switch v := curr.(type) {
		case int:
			res = append(res, v)
		case *TreeNode:
			stack.Push(v.value)
			if v.right != nil {
				stack.Push(v.right)
			}
			if v.left != nil {
				stack.Push(v.left)
			}
		}

	}

	return res
}

func (t *BinarySearchTree) PreOrder() []int {
	res := []int{}
	stack := NewStack()

	stack.Push(t.root)
	for stack.Len() > 0 {
		curr := stack.Pop()
		switch v := curr.(type) {
		case int:
			res = append(res, v)
		case *TreeNode:
			if v.right != nil {
				stack.Push(v.right)
			}
			if v.left != nil {
				stack.Push(v.left)
			}
			stack.Push(v.value)
		}

	}

	return res
}

func (t *BinarySearchTree) Inorder() []int {
	res := []int{}
	stack := NewStack()

	stack.Push(t.root)
	for stack.Len() > 0 {
		curr := stack.Pop()
		switch v := curr.(type) {
		case int:
			res = append(res, v)
		case *TreeNode:
			if v.right != nil {
				stack.Push(v.right)
			}
			stack.Push(v.value)
			if v.left != nil {
				stack.Push(v.left)
			}
		}

	}

	return res
}
