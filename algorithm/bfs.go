package algorithm

import (
	"fmt"
	"interview_go/datastructure"
)

func BfsIteration(graph *datastructure.Graph, value any) {
	if graph == nil {
		return
	}

	queue := []*datastructure.Graph{graph}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if !curr.Visited {
			fmt.Println(curr.Value)
			curr.Visited = true
			if curr.Neighbors == nil {
				continue
			}
			for _, n := range curr.Neighbors {
				if !n.Visited {
					queue = append(queue, n)
				}
			}
		}
	}
}

type RowCol struct {
	row int
	col int
}

func FloodFill_BfsExample(img [][]int, row, col, p int) [][]int {

	start := img[row][col]
	queue := []RowCol{{row, col}}
	visited := make(map[RowCol]bool)

	for len(queue) > 0 {
		rc := queue[0]
		queue = queue[1:]
		visited[rc] = true
		img[rc.row][rc.col] = p

		for _, nrc := range FloodFill_FindNeighbors(img, rc, start) {
			if _, ok := visited[nrc]; !ok {
				queue = append(queue, nrc)
			}
		}
	}

	return img
}

func FloodFill_FindNeighbors(img [][]int, rc RowCol, start int) []RowCol {
	result := []RowCol{}

	if rc.row < 0 || rc.row >= len(img) || rc.col < 0 || rc.col >= len(img[0]) {
		return result
	}

	// left
	if rc.col-1 >= 0 && img[rc.row][rc.col-1] == start {
		result = append(result, RowCol{row: rc.row, col: rc.col - 1})
	}
	// right
	if rc.col+1 < len(img[rc.row]) && img[rc.row][rc.col+1] == start {
		result = append(result, RowCol{row: rc.row, col: rc.col + 1})
	}
	// up
	if rc.row-1 >= 0 && img[rc.row-1][rc.col] == start {
		result = append(result, RowCol{row: rc.row - 1, col: rc.col})
	}
	// down
	if rc.row+1 < len(img) && img[rc.row+1][rc.col] == start {
		result = append(result, RowCol{row: rc.row + 1, col: rc.col})
	}

	return result
}
