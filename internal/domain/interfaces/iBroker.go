package interfaces

type IBroker interface {
	Publish(topic string, message any)
}
