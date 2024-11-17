package network

import (
	"fmt"
)

// var n Node = Node{5, make([]*Node, 0, 5)}
type hypercubeNode struct {
	id            string
	neighbors     []*hypercubeNode
	neighborCount int
	inputQ        chan string
}

func (n *hypercubeNode) GetId() string {
	return n.id
}

func (n *hypercubeNode) AcceptMessage(msg string) {
	n.inputQ <- msg
}

func (n *hypercubeNode) Start(master Master, finishedMsg string) {

	go func() {

		for {

			select {
			case fromNode := <-n.inputQ:

				switch fromNode {
				case "-1":

					fmt.Println(n.id, " initiaiting conversation...")

					for nbr := 0; nbr < n.neighborCount; nbr++ {

						n.neighbors[nbr].inputQ <- n.id

						fromNode = <-n.inputQ

						fmt.Println(n.id, "heard back from", fromNode+".")
					}

					master.AcceptMessage(finishedMsg)

				default:

					fmt.Println(n.id, " heard from ", fromNode+". Responding.")

					for nbr := 0; nbr < n.neighborCount; nbr++ {

						if n.neighbors[nbr].id == fromNode {

							n.neighbors[nbr].inputQ <- n.id

						}

					}

				}
			}

		}

	}()

}
