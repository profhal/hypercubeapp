package network

import (
	"fmt"
)

// var n Node = Node{5, make([]*Node, 0, 5)}
type gridNode struct {
	id            string
	right         *gridNode
	up            *gridNode
	left          *gridNode
	down          *gridNode
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
				case "start":

					fmt.Println(n.id, " initiaiting conversation...")

					n.right.inputQ <- n.id
					fromNode = <-n.inputQ
					fmt.Println(n.id, "heard back from", fromNode+".")

					n.up.inputQ <- n.id
					fromNode = <-n.inputQ
					fmt.Println(n.id, "heard back from", fromNode+".")

					n.left.inputQ <- n.id
					fromNode = <-n.inputQ
					fmt.Println(n.id, "heard back from", fromNode+".")

					n.down.inputQ <- n.id
					fromNode = <-n.inputQ
					fmt.Println(n.id, "heard back from", fromNode+".")

					master.NodeFinished()

				default:

					fmt.Println(n.id, " heard from ", fromNode+". Responding.")

					switch fromNode {
					case n.right.id:
						n.right.inputQ <- n.id
					case n.up.id:
						n.up.inputQ <- n.id
					case n.left.id:
						n.left.inputQ <- n.id
					case n.down.id:
						n.down.inputQ <- n.id
					}
				}
			}

		}

	}()

}
