package hypercube

import (
	"fmt"
	"strconv"
)

// var n Node = Node{5, make([]*Node, 0, 5)}
type node struct {
	id        int
	dimension int
	neighbors []*node
	inputQ    chan int
}

func (n *node) start(master Master, finishedValue int) {

	go func() {

		for {

			select {
			case fromNode := <-n.inputQ:

				switch fromNode {
				case -1:

					fmt.Println(n.id, " initiaiting conversation...")

					for nbr := 0; nbr < n.dimension; nbr++ {

						n.neighbors[nbr].inputQ <- n.id

						fromNode = <-n.inputQ

						fmt.Println(n.id, "heard back from", strconv.Itoa(fromNode)+".")
					}

					master.AcceptMessage(finishedValue)

				default:

					fmt.Println(n.id, " heard from ", strconv.Itoa(fromNode)+". Responding.")

					for nbr := 0; nbr < n.dimension; nbr++ {

						if n.neighbors[nbr].id == fromNode {

							n.neighbors[nbr].inputQ <- n.id

						}

					}

				}
			}

		}

	}()

}
