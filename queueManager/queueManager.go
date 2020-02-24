package queueManager

type QueueManager interface {
	Subscribe(topic string, cb func(data interface{}))
	Publish(topic string, payload interface{}) error
}
