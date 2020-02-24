package queueManager

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	defaultTimeout = time.Second
	defaultQOS     = 0
)

var log = logrus.StandardLogger()

type mqttQueueManager struct {
	client mqtt.Client
}

func MakeMQTTQueueManager(dsn string) (QueueManager, error) {
	d := &mqttQueueManager{}
	err := d.connect(dsn)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (q *mqttQueueManager) Subscribe(topic string, cb func(data interface{})) {
	q.client.Subscribe(topic, defaultQOS, func(client mqtt.Client, message mqtt.Message) {
		log.Debugf("Received message in topic %s", message.Topic())
		cb(message.Payload())
	})
}

func (q *mqttQueueManager) Publish(topic string, payload interface{}) error {
	if q.client != nil {
		t := q.client.Publish(topic, defaultQOS, true, payload)
		if !t.WaitTimeout(defaultTimeout) {
			return fmt.Errorf("timeout publishing")
		}
		return t.Error()
	}

	return fmt.Errorf("not connected")
}

func (q *mqttQueueManager) connect(dsn string) error {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(dsn)

	opts.OnConnect = q.onConnect

	c := mqtt.NewClient(opts)
	c.Connect().WaitTimeout(defaultTimeout)

	if !c.IsConnected() {
		return fmt.Errorf("error connecting to %s: timeout", dsn)
	}

	q.client = c

	return nil
}

func (q *mqttQueueManager) onConnect(client mqtt.Client) {
	log.Debug("Connected to MQTT")
}
