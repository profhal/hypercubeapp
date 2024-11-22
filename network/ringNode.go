package network

import (
	"fmt"
	"strconv"
)

type ringNode struct {
	Node
	id            string
	left          *ringNode
	right         *ringNode
	neighborCount int
	inputQ        chan message
}

func (n *ringNode) GetId() string {
	return n.id
}

func (n *ringNode) AcceptMessage(msg message) {
	n.inputQ <- msg
}

func (n *ringNode) Start(master Master, finishedMsg string) {

	go func() {

		for {

			select {
			case msg := <-n.inputQ:

				switch msg.message {
				case "left":

					switch msg.senderId {
					case NETWORK_MASTER:

						fmt.Println(n.id, "starting left-wise loop...")

						n.left.AcceptMessage(message{"left", n.id})

						<-n.inputQ

						master.NodeFinished()

					default:

						fmt.Println(n.id, "heard from", msg.senderId)

						n.left.AcceptMessage(message{"left", n.id})

					}

				case "right":

					switch msg.senderId {
					case NETWORK_MASTER:

						fmt.Println(n.id, "starting right-wise loop...")

						n.right.AcceptMessage(message{"right", n.id})

						<-n.inputQ

						master.NodeFinished()

					default:

						n.right.AcceptMessage(message{"right", n.id})

					}

				case "elect":

					switch msg.senderId {
					case NETWORK_MASTER:

						fmt.Println("Starting election...")

						n.right.AcceptMessage(message{"elect", n.id})

					default:

						if msg.senderId == n.id {

							fmt.Println("The ring has elected: ", n.id)

							fmt.Println()

							master.NodeFinished()

						} else {

							senderId, _ := strconv.Atoi(msg.senderId)
							localId, _ := strconv.Atoi(n.id)

							if senderId < localId {

								n.right.AcceptMessage(message{"elect", n.id})

							} else {

								n.right.AcceptMessage(message{"elect", msg.senderId})

							}

						}

					}

				}

			}

		}

	}()

}
