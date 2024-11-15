package hypercube

import (
	"math"
)

// A dimension-D hypercube has 2^{dimesion} nodes. So, nodeCount = 2^{dimension} and the nodes
// slice is nodeCount long once the hypercube is initialized.
type Hypercube struct {
	dimension int
	nodeCount int
	nodes     []*node
}

// Returns a pointer to a dimension-D hypercube.
func CreateHypercube(dimension int) *Hypercube {

	// Prep the hypercube.
	//
	hypercube := new(Hypercube)

	hypercube.dimension = dimension
	hypercube.nodeCount = int(math.Pow(2, float64(dimension)))

	hypercube.nodes = make([]*node, 0, hypercube.nodeCount)

	for n := 0; n < hypercube.nodeCount; n++ {

		hypercube.nodes = append(hypercube.nodes, new(node))

		hypercube.nodes[n].id = n
		hypercube.nodes[n].dimension = hypercube.dimension
		hypercube.nodes[n].neighbors = make([]*node, 0, dimension)
		hypercube.nodes[n].inputQ = make(chan int)

		hypercube.nodes[n].start()

	}

	// Wire hypercube together.
	//
	powersOfTwo := []int{}

	for d := 0; d < hypercube.dimension; d++ {

		powersOfTwo = append(powersOfTwo, int(math.Pow(float64(2), float64(d))))

	}

	for n := 0; n < hypercube.nodeCount; n++ {

		for p := 0; p < hypercube.dimension; p++ {

			hypercube.nodes[n].neighbors = append(hypercube.nodes[n].neighbors, hypercube.nodes[(n^powersOfTwo[p])])

		}

	}

	return hypercube

}

// Runs the hypercube task.
//
// At the moment this exists so that we can time the creation and investigate
// memory use. That is, we need something public to reference so that we can
// use a hypercube object to satisfy compilation.
func (h *Hypercube) Run() {

	h.nodes[0].inputQ <- -1

}
