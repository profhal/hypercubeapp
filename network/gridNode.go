package network

import (
	"fmt"
)

type gridNode struct {
	id            string
	right         *gridNode
	up            *gridNode
	left          *gridNode
	down          *gridNode
	neighborCount int
	inputQ        chan message
}

func (n *gridNode) GetId() string {
	return n.id
}

func (n *gridNode) AcceptMessage(msg message) {
	n.inputQ <- msg
}

func (n *gridNode) Start(master Master, finishedMsg string) {

	go func() {

		for {

			select {
			case msg := <-n.inputQ:

				switch msg.message {
				case "start":

					if msg.senderId == NETWORK_MASTER {

						fmt.Println(n.id, " initiaiting conversation...")

						if n.right != nil {
							n.right.AcceptMessage(message{"hello", n.id})
							responseMsg := <-n.inputQ
							fmt.Println(n.id, "heard back from", responseMsg.senderId+".")
						}

						if n.up != nil {
							n.up.AcceptMessage(message{"hello", n.id})
							responseMsg := <-n.inputQ
							fmt.Println(n.id, "heard back from", responseMsg.senderId+".")
						}

						if n.left != nil {
							n.left.AcceptMessage(message{"hello", n.id})
							responseMsg := <-n.inputQ
							fmt.Println(n.id, "heard back from", responseMsg.senderId+".")
						}

						if n.down != nil {
							n.down.AcceptMessage(message{"hello", n.id})
							responseMsg := <-n.inputQ
							fmt.Println(n.id, "heard back from", responseMsg.senderId+".")
						}

						master.NodeFinished()

					}

				case "hello":

					fmt.Println(n.id, " heard from ", msg.senderId+". Responding.")

					if n.right != nil && n.right.id == msg.senderId {

						n.right.AcceptMessage(message{"", n.id})

					} else if n.up != nil && n.up.id == msg.senderId {

						n.up.AcceptMessage(message{"", n.id})

					} else if n.left != nil && n.left.id == msg.senderId {

						n.left.AcceptMessage(message{"", n.id})

					} else if n.down != nil && n.down.id == msg.senderId {

						n.down.AcceptMessage(message{"", n.id})

					}

				}

			}

		}

	}()

}
