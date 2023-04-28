package mock

type BrokerFake struct {
}

func NewBrokerFake() *BrokerFake {
	return &BrokerFake{}
}

func (b *BrokerFake) Publish(topic string, message any) {

}
