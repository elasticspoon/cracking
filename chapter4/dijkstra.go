package chapter4

import (
	"container/heap"
	"math"
)

type DNode struct {
	value int
	dist  int
}

type PriorityQueue []*DNode

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*DNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Dijkstra(graph map[int][]*DNode, start int) []int {
	distances := make(map[int]*DNode)
	for k := range graph {
		distances[k] = &DNode{value: k, dist: math.MaxInt}
	}

	distances[start] = &DNode{value: start, dist: 0}
	visited := make(map[int]bool)
	result := []int{}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*DNode).value
		if _, ok := visited[curr]; ok {
			continue
		}

		visited[curr] = true
		result = append(result, curr)

		for _, n := range graph[curr] {
			currDist := distances[curr].dist
			potentialDist := currDist + n.dist
			if potentialDist > currDist {
				distances[n.value].dist = potentialDist
				heap.Push(&pq, n)
			}
		}

	}

	return result
}
