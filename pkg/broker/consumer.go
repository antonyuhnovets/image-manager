package broker

import (
	"fmt"
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

func (c *Consumer) Connect(cfg *config.Config) {
	c.Connection = helpers.ConnectAMQP(cfg.AMQP_URL)
	c.Channel = helpers.OpenChannel(c.Connection)
	c.Queue = helpers.DeclareQueue("img", c.Channel)
	log.Println("Consumer connectd")
	c.Storage = storage.SetStorage(cfg)
	log.Println("Consumer storage set")
}

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
			fmt.Println("consuming")
			time.Sleep(time.Second * 5)
			controllers.CompressAndSave(c.Storage, msg.Body)
			fmt.Println("Accepted img")
		}
	}()
	<-forever
}

func (c *Consumer) Disconnect() {
	c.Connection.Close()
	c.Channel.Close()
}
