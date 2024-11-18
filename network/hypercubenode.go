package network

import (
	"fmt"
)

type hypercubeNode struct {
	id            string
	neighbors     []*hypercubeNode
	neighborCount int
	inputQ        chan message
}

func (n *hypercubeNode) GetId() string {
	return n.id
}

func (n *hypercubeNode) AcceptMessage(msg message) {
	n.inputQ <- msg
}

func (n *hypercubeNode) Start(master Master, finishedMsg string) {

	go func() {

		hellos := 0

		for {

			select {
			case msg := <-n.inputQ:

				switch msg.message {
				case "contact neighbors":

					fmt.Println(n.id, " initiaiting conversation...")

					hellos = n.neighborCount

					for nbr := 0; nbr < n.neighborCount; nbr++ {

						n.neighbors[nbr].AcceptMessage(message{"hello", n.id})

					}

				case "hello":

					if hellos > 0 {

						fmt.Println(n.id, " heard back from ", msg.senderId+".")

						hellos--

						if hellos == 0 {

							master.NodeFinished()

						}

					} else {

						fmt.Println(n.id, " heard from ", msg.senderId+". Responding.")

						for nbr := 0; nbr < n.neighborCount; nbr++ {

							if n.neighbors[nbr].id == msg.senderId {

								n.neighbors[nbr].AcceptMessage(message{"hello", n.id})

							}

						}

					}

				}

			}

		}

	}()

}
