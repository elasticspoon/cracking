package chapter4

import (
	"math/rand"
)

type BinTree struct {
	left     *BinTree
	right    *BinTree
	numNodes int
	value    int
}

func (b *BinTree) RandomNode() *BinTree {
	randVal := rand.Intn(b.numNodes) + 1
	visited := make(map[*BinTree]bool)
	stack := []*BinTree{b}

	for count, curr := 0, b; curr != nil; {
		// current is processed
		// fmt.Println("curr", curr.value, stack)
		if visited[curr] {

			// fmt.Println("visited")
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			continue
		}

		// has left child and not visited go deeper left
		if curr.left != nil && !visited[curr.left] {
			stack = append(stack, curr.left)
			curr = curr.left
			continue
		}

		// process current
		count++
		// fmt.Println("processing", count, curr.value)
		if count == randVal {
			return curr
		}
		visited[curr] = true

		if curr.right != nil && !visited[curr.right] {
			stack = append(stack, curr.right)
			curr = curr.right
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

	}
	panic("Shouldn't get here")
}

func (t *BinTree) NumPathsWithSum(tar int) int {
	count := t.recNumPathsWithSum(tar, 0)
	if t.left != nil {
		count += t.left.NumPathsWithSum(tar)
	}

	if t.right != nil {
		count += t.right.NumPathsWithSum(tar)
	}

	return count
}

func (t *BinTree) recNumPathsWithSum(tar int, sum int) int {
	count := 0

	if t.value+sum == tar {
		count++
	}

	if t.left != nil {
		count += t.left.recNumPathsWithSum(tar, t.value+sum)
	}
	if t.right != nil {
		count += t.right.recNumPathsWithSum(tar, t.value+sum)
	}
	return count
}
