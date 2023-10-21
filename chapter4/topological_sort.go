package chapter3

type NodeTop struct {
	dep   []int
	preq  int
	value int
}

func TopologicalSort(input [][2]int) []int {
	graph := make(map[int]*NodeTop)

	for _, edge := range input {
		s, e := edge[0], edge[1]

		start, ok := graph[s]
		if ok {
			start.dep = append(start.dep, e)
		} else {
			graph[s] = &NodeTop{dep: []int{e}, value: s}
		}

		end, ok := graph[e]
		if ok {
			end.preq++
		} else {
			graph[e] = &NodeTop{dep: []int{}, preq: 1, value: e}
		}
	}

	result := []int{}
	queue := []*NodeTop{}

	for _, node := range graph {
		if node.preq == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current.value)

		for _, n := range current.dep {
			node := graph[n]
			node.preq--
			if node.preq <= 0 {
				queue = append(queue, node)
			}
		}
	}

	return result
}
