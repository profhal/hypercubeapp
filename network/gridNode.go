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

					if n.right != nil {
						n.right.inputQ <- n.id
						fromNode = <-n.inputQ
						fmt.Println(n.id, "heard back from", fromNode+".")
					}

					if n.up != nil {
						n.up.inputQ <- n.id
						fromNode = <-n.inputQ
						fmt.Println(n.id, "heard back from", fromNode+".")
					}

					if n.left != nil {
						n.left.inputQ <- n.id
						fromNode = <-n.inputQ
						fmt.Println(n.id, "heard back from", fromNode+".")
					}

					if n.down != nil {
						n.down.inputQ <- n.id
						fromNode = <-n.inputQ
						fmt.Println(n.id, "heard back from", fromNode+".")
					}
					master.NodeFinished()

				default:

					fmt.Println(n.id, " heard from ", fromNode+". Responding.")

					if n.right != nil && n.right.id == fromNode {
						n.right.inputQ <- n.id
					} else if n.up != nil && n.up.id == fromNode {
						n.up.inputQ <- n.id
					} else if n.left != nil && n.left.id == fromNode {
						n.left.inputQ <- n.id
					} else if n.down != nil && n.down.id == fromNode {
						n.down.inputQ <- n.id
					}

				}

			}

		}

	}()

}
