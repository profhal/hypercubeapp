package network

import (
	"math"
	"strconv"
)

// A dimension-D hypercube has 2^{dimesion} nodes. So, nodeCount = 2^{dimension} and the nodes
// slice is nodeCount long once the hypercube is initialized.
type Hypercube struct {
	Master
	dimension int
	nodeCount int
	nodes     []*hypercubeNode
	inputQ    chan string
}

// Returns a pointer to a dimension-D hypercube.
func CreateHypercube(dimension int) *Hypercube {

	// Prep the hypercube.
	//
	hypercube := new(Hypercube)

	hypercube.dimension = dimension
	hypercube.nodeCount = int(math.Pow(2, float64(dimension)))
	hypercube.inputQ = make(chan string)

	hypercube.nodes = make([]*hypercubeNode, 0, hypercube.nodeCount)

	for n := 0; n < hypercube.nodeCount; n++ {

		hypercube.nodes = append(hypercube.nodes, new(hypercubeNode))

		hypercube.nodes[n].id = strconv.Itoa(n)
		hypercube.nodes[n].neighborCount = hypercube.dimension
		hypercube.nodes[n].neighbors = make([]*hypercubeNode, 0, hypercube.dimension)
		hypercube.nodes[n].inputQ = make(chan string, hypercube.dimension)

		hypercube.nodes[n].Start(hypercube, "0")

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
func (h *Hypercube) Touch(nodeId int) {

	h.nodes[nodeId].inputQ <- "-1"

	<-h.inputQ

}

func (h *Hypercube) AcceptMessage(msg string) {
	h.inputQ <- msg
}
