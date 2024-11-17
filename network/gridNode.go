package network

import (
	"fmt"
)

const (
	grid_right     = 0
	grid_up    int = 1
	grid_left  int = 2
	grid_down  int = 3
)

// var n Node = Node{5, make([]*Node, 0, 5)}
type gridNode struct {
	id            string
	neighbors     []*gridNode
	neighborCount int
	inputQ        chan string
}

func (n *gridNode) GetId() string {
	return n.id
}

func (n *gridNode) AcceptMessage(msg string) {
	n.inputQ <- msg
}

func (n *gridNode) Start(master Master, finishedMsg string) {

	go func() {

		for {

			select {
			case fromNode := <-n.inputQ:

				switch fromNode {
				case "-1":

					fmt.Println(n.id, " initiaiting conversation...")

					n.neighbors[grid_right].inputQ <- n.id
					fromNode = <-n.inputQ
					fmt.Println(n.id, "heard back from", fromNode+".")

					n.neighbors[grid_up].inputQ <- n.id
					fromNode = <-n.inputQ
					fmt.Println(n.id, "heard back from", fromNode+".")

					n.neighbors[grid_left].inputQ <- n.id
					fromNode = <-n.inputQ
					fmt.Println(n.id, "heard back from", fromNode+".")

					n.neighbors[grid_down].inputQ <- n.id
					fromNode = <-n.inputQ
					fmt.Println(n.id, "heard back from", fromNode+".")

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
