package broker

type IEventBus interface {
	Publish(string, interface{})
	Subscribe(string, DataChannel)
}
