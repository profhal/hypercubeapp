package network

type Node interface {
	GetId() string
	AddMessage(msg string)
	Start(master *Master, finishedMsg string)
}
