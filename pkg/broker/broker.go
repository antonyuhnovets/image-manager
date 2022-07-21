// Abstract interfaces for message broker

package broker

import (
	"log"
	"mime/multipart"

	"github.com/antonyuhnovets/image-manager/pkg/config"
)

// Producer (publisher) abstract interface
// Publish method to put message (file in this case) in queue
type Producer interface {
	Publish(string, multipart.File, *multipart.FileHeader) error
	Disconnect()
}

// Consumer have Consume method to start accepting messages from queue
// Takes handler function for messages as argument
type Consumer interface {
	Consume(func(string, []byte) error)
	Disconnect()
}

// Put together parts of msg broker into single essence
type Broker struct {
	Producer Producer
	Consumer Consumer
}

// Connect producer and consumer, pass to broker and return
func SetBroker(cfg *config.Config) *Broker {
	p := ConnectRMQ(cfg)
	log.Println("Producer connected")
	c := ConnectRMQ(cfg)
	log.Println("Consumer connected")
	broker := &Broker{
		p,
		c,
	}

	return broker
}
