package datastructure

type Graph struct {
	Value     any
	Neighbors []*Graph
	Visited   bool
}
