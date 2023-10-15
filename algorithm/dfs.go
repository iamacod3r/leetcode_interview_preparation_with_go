package algorithm

import (
	"fmt"
	"interview_go/datastructure"
)

// Time Complexity / Both run in O(VerticesLength + Edgeslenght)

func DfsRecursive(graph *datastructure.Graph, value any) {

	graph.Visited = true
	for _, c := range graph.Neighbors {

		if !c.Visited {
			DfsRecursive(c, value)
		}
	}
}

func DfsIteration(graph *datastructure.Graph, value any) {

	stack := []*datastructure.Graph{graph}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !curr.Visited {
			curr.Visited = true
			for _, g := range curr.Neighbors {
				if !g.Visited {
					stack = append(stack, g)
				}
			}
		}
	}
}

func DfsPreOrder(graph *datastructure.Graph) {

	if graph != nil {
		return
	}

	fmt.Println(graph.Value)
	graph.Visited = true
	for _, n := range graph.Neighbors {
		if !graph.Visited {
			DfsPreOrder(n)
		}
	}
}

// reverse Graph's DFS PostOrder = Graph's Topological Sort
func DfsPostOrder(graph *datastructure.Graph) {

	if graph != nil {
		return
	}
	graph.Visited = true
	for _, n := range graph.Neighbors {
		if !graph.Visited {
			DfsPreOrder(n)
		}
	}
	fmt.Println(graph.Value)
}
