package hypercube

// var n Node = Node{5, make([]*Node, 0, 5)}
type node struct {
	dimension int

	neighbors []*node
}
