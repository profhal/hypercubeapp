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
			case msg := <-n.inputQ:

				switch msg {
				case "start":

					fmt.Println(n.id, " initiaiting conversation...")

					for nbr := 0; nbr < n.neighborCount; nbr++ {

						n.neighbors[nbr].inputQ <- n.id

						msg = <-n.inputQ

						fmt.Println(n.id, "heard back from", msg+".")
					}

					master.NodeFinished()

				default:

					fmt.Println(n.id, " heard from ", msg+". Responding.")

					for nbr := 0; nbr < n.neighborCount; nbr++ {

						if n.neighbors[nbr].id == msg {

							n.neighbors[nbr].inputQ <- n.id

						}

					}

				}
			}

		}

	}()

}
