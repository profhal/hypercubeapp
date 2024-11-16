package hypercube

type Master interface {
	AcceptMessage(msg int)
}
