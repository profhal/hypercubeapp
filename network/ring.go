package network

import (
	"strconv"
)

type Ring struct {
	Master
	nodeCount int
	nodes     []*ringNode
	inputQ    chan string
}

func CreateRing(nodeCount int) *Ring {

	ring := new(Ring)

	// Prep the nodes
	//
	ring.nodeCount = nodeCount
	ring.inputQ = make(chan string)

	ring.nodes = make([]*ringNode, 0, ring.nodeCount)

	for n := 0; n < ring.nodeCount; n++ {

		ring.nodes = append(ring.nodes, new(ringNode))

		ring.nodes[n].id = strconv.Itoa(n)

		ring.nodes[n].neighborCount = 2
		ring.nodes[n].neighbors = make([]*ringNode, 2)
		ring.nodes[n].inputQ = make(chan string, 2)

		ring.nodes[n].Start(ring, "0")

	}

	// Wire the ring
	//
	for n := range ring.nodes {

		if n == 0 {

			ring.nodes[n].neighbors[ring_left] = ring.nodes[ring.nodeCount-1]
			ring.nodes[n].neighbors[ring_right] = ring.nodes[n+1]

		} else if n == ring.nodeCount-1 {

			ring.nodes[n].neighbors[ring_left] = ring.nodes[n-1]
			ring.nodes[n].neighbors[ring_right] = ring.nodes[0]

		} else {

			ring.nodes[n].neighbors[ring_left] = ring.nodes[n-1]
			ring.nodes[n].neighbors[ring_right] = ring.nodes[n+1]

		}

	}

	return ring

}

// Runs the grid task
func (r *Ring) Loop(startWith int, direction string) {

	r.nodes[startWith].AcceptMessage(direction)

	<-r.inputQ

}

func (r *Ring) AcceptMessage(msg string) {
	r.inputQ <- msg
}
