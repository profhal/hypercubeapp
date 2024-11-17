package network

import (
	"fmt"
	"strings"
)

// var n Node = Node{5, make([]*Node, 0, 5)}
type ringNode struct {
	id            string
	left          *ringNode
	right         *ringNode
	neighborCount int
	inputQ        chan string
}

func (n *ringNode) GetId() string {
	return n.id
}

func (n *ringNode) AcceptMessage(msg string) {
	n.inputQ <- msg
}

func (n *ringNode) Start(master Master, finishedMsg string) {

	go func() {

		for {

			select {
			case msg := <-n.inputQ:

				switch msg {
				case "left":

					fmt.Println(n.id, "starting left-wise loop...")

					n.left.inputQ <- "left " + n.id

					<-n.inputQ

					master.NodeFinished()

				case "right":

					fmt.Println(n.id, "starting right-wise loop...")

					n.right.inputQ <- "right " + n.id

					<-n.inputQ

					master.NodeFinished()

				default:

					// If it's not just "left" or "right" it will be "left x" or right x" where x is the
					// node sending the message. This means we should pass left - or right - and not
					// wait for a response.
					directionAndSender := strings.Split(msg, " ")

					fmt.Println(n.id, " heard from", directionAndSender[1]+". Responding.")

					if directionAndSender[0] == "left" {

						n.left.inputQ <- "left " + n.id

					} else {

						n.right.inputQ <- "right " + n.id

					}

				}
			}

		}

	}()

}
