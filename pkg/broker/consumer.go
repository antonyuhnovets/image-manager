package broker

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/antonyuhnovets/image-manager/pkg/config"
	"github.com/antonyuhnovets/image-manager/pkg/controllers"
	"github.com/antonyuhnovets/image-manager/pkg/helpers"
	"github.com/antonyuhnovets/image-manager/pkg/storage"
)

type Consumer struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
	Storage    storage.Entity
}

// Connect consumer, create channel, declare queue and get storage
func (c *Consumer) Connect(cfg *config.Config) {
	c.Connection = helpers.ConnectAMQP(cfg.AMQP_URL)
	c.Channel = helpers.OpenChannel(c.Connection)
	c.Queue = helpers.DeclareQueue("img", c.Channel)
	log.Println("Consumer connectd")

	c.Storage = storage.GetStorage(cfg)
	log.Println("Consumer storage set")
}

// Declare Qos for message buffer
// Start consuming messages and handle it
func (c *Consumer) Consume() {
	err := c.Channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Fatalf("Failed to set QoS: %s", err)
	}
	msgs, err := c.Channel.Consume(
		c.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error consuming messages: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			log.Println("Consuming")
			time.Sleep(time.Second * 5)
			err = controllers.CompressAndSave(msg.MessageId, c.Storage, msg.Body)
			if err != nil {
				log.Println(err)
			}
		}
	}()
	<-forever
}

func (c *Consumer) Disconnect() {
	c.Connection.Close()
	c.Channel.Close()
}
