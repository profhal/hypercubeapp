package network

import (
	"fmt"
)

// var n Node = Node{5, make([]*Node, 0, 5)}
type node struct {
	id            string
	neighbors     []*node
	neighborCount int
	inputQ        chan string
}

func (n *node) start(master Master, finishedValue string) {

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

					master.AcceptMessage(finishedValue)

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
