package network

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Ring struct {
	Master
	nodeCount int
	nodes     []*ringNode
	inputQ    chan string
}

func CreateRing(nodeCount int) *Ring {

	isUsed := make(map[int]bool)

	ring := new(Ring)

	// Prep the nodes
	//
	ring.nodeCount = nodeCount
	ring.inputQ = make(chan string)

	ring.nodes = make([]*ringNode, 0, ring.nodeCount)

	for n := 0; n < ring.nodeCount; n++ {

		ring.nodes = append(ring.nodes, new(ringNode))

		possibleId := rand.Intn(10 * ring.nodeCount)

		_, ok := isUsed[possibleId]

		for ok {

			possibleId := rand.Intn(10 * ring.nodeCount)

			_, ok = isUsed[possibleId]

		}

		isUsed[possibleId] = true

		ring.nodes[n].id = strconv.Itoa(possibleId)

		ring.nodes[n].neighborCount = 2
		ring.nodes[n].inputQ = make(chan message, 2)

		ring.nodes[n].Start(ring, "0")

	}

	// Wire the ring
	//
	for n := range ring.nodes {

		if n == 0 {

			ring.nodes[n].left = ring.nodes[ring.nodeCount-1]
			ring.nodes[n].right = ring.nodes[n+1]

		} else if n == ring.nodeCount-1 {

			ring.nodes[n].left = ring.nodes[n-1]
			ring.nodes[n].right = ring.nodes[0]

		} else {

			ring.nodes[n].left = ring.nodes[n-1]
			ring.nodes[n].right = ring.nodes[n+1]

		}

	}

	return ring

}

// Runs the grid task
func (r *Ring) Loop(startWith int, direction string) {

	r.nodes[startWith].AcceptMessage(message{direction, NETWORK_MASTER})

	<-r.inputQ

}

func (r *Ring) RunElection() {

	fmt.Println("Node IDs:")

	for n := range r.nodes {
		fmt.Println("   ", r.nodes[n].GetId())
	}

	r.nodes[0].AcceptMessage(message{"elect", NETWORK_MASTER})

	<-r.inputQ

}

func (r *Ring) NodeFinished() {
	r.inputQ <- "done"
}
