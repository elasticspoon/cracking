package chapter3

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T) {
	g := map[int][][2]int{
		0: {{1, 7}, {2, 9}, {5, 14}},
		1: {{0, 7}, {2, 10}, {3, 15}},
		2: {{0, 9}, {1, 10}, {3, 11}, {5, 2}},
		3: {{1, 15}, {2, 11}, {4, 6}},
		4: {{3, 6}, {5, 9}},
		5: {{0, 14}, {2, 2}, {4, 9}},
	}

	graph := make(map[int][]*DNode)
	for k, v := range g {
		graph[k] = make([]*DNode, len(v))
		for i, n := range v {
			graph[k][i] = &DNode{value: n[0], dist: n[1]}
		}
	}

	start := 0
	expected := []int{0, 7, 9, 20, 20, 11}
	path := Dijkstra(graph, start)

	for i, v := range path {
		if v != expected[i] {
			t.Errorf("Expected %v, got %v", expected, path)
			break
		}
	}
}

func showGraph(g map[int][]*DNode) string {
	var out bytes.Buffer

	out.WriteString("\n")
	for k, v := range g {
		out.WriteString(fmt.Sprint("k: ", k, " v: "))
		for _, n := range v {
			out.WriteString(fmt.Sprintf("(%d, %d)", n.value, n.dist))
		}
		out.WriteString("\n")
	}
	return out.String()
}
