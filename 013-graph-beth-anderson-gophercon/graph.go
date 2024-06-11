package graph

type Node[T comparable] struct {
	Value T
	Edges []*Edge[T]
}

func NewNode[T comparable](val T) Node[T] {
	return Node[T]{Value: val}
}

func (n *Node[T]) AddEdge(to *Node[T], weight int) {
	n.Edges = append(n.Edges, &Edge[T]{ToNode: to, Weight: weight})
}

type Edge[T comparable] struct {
	ToNode  *Node[T]
	Weight  int
	visited bool
}

func unvisited[T comparable](edges []*Edge[T]) []*Node[T] {
	unv := []*Node[T]{}
	for _, edge := range edges {
		if !edge.visited {
			unv = append(unv, edge.ToNode)
			edge.visited = true
		}
	}
	return unv
}

func bfs[T comparable](node *Node[T], vals *[]T) {
	if node != nil {
		*vals = append(*vals, node.Value)
		for _, n := range unvisited(node.Edges) {
			bfs(n, vals)
		}
	}
}

func BreadthFirst[T comparable](start *Node[T]) []T {
	vals := []T{}
	bfs(start, &vals)
	return vals
}
