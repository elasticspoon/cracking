package ctci

import (
	"math"
)

type LLNode[V any] struct {
	value V
	next  *LLNode[V]
}

func genRevNum(head *LLNode[int]) int {
	num := 0
	for head != nil {
		num = num*10 + head.value
		head = head.next
	}
	return num
}

func genFwdNum(head *LLNode[int]) int {
	length, num := 0, 0

	for curr := head; curr != nil; curr = curr.next {
		length++
	}

	for curr, i := head, length-1; curr != nil; curr, i = curr.next, i-1 {
		num += curr.value * int(math.Pow10(i))
	}
	return num
}

func AddListsFwd(a, b *LLNode[int]) int {
	return genFwdNum(a) + genFwdNum(b)
}

func AddListsRev(a, b *LLNode[int]) int {
	return genRevNum(a) + genRevNum(b)
}

func ListIsPal(head *LLNode[any]) bool {
	rev := revNodeList(head)

	for ; head != nil; head, rev = head.next, rev.next {
		if head.value != rev.value {
			return false
		}
	}

	return true
}

func revNodeList(head *LLNode[any]) *LLNode[any] {
	var newHead *LLNode[any]
	for ; head != nil; head = head.next {
		newHead = &LLNode[any]{value: head.value, next: newHead}
	}

	return newHead
}

func FindCycle(head *LLNode[any]) *LLNode[any] {
	fast := head.next.next
	slow := head.next

	for fast != slow {
		fast = fast.next.next
		slow = slow.next
	}

	return slow
}
