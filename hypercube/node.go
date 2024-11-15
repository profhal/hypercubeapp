package hypercube

import (
	"fmt"
	"time"
)

// var n Node = Node{5, make([]*Node, 0, 5)}
type node struct {
	id        int
	dimension int
	neighbors []*node
	inputQ    chan int
}

func (n *node) start() {

	go func() {

		for {

			select {
			case fromNode := <-n.inputQ:

				fmt.Println(n.id, "heard from", fromNode)

				for nbr := 0; nbr < n.dimension; nbr++ {

					n.neighbors[nbr].inputQ <- n.id

					time.Sleep(10 * time.Millisecond)

				}

			}

		}

	}()

}
