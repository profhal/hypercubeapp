package network

type Master interface {
	AcceptMessage(msg string)
}
