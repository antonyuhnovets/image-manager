// Realizaton broker methods for RabbitMQ

package broker

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/antonyuhnovets/image-manager/pkg/config"
)

// Struct with set of required essences for connecting to RMQ server
// Used to setup producer and consumer in the same way
type RMQConnection struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

// Dial RabbitMQ by connection string
func ConnectAMQP(url string) *amqp.Connection {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect RabbitMQ: %s", err)
	}

	return conn
}

// Return opened channel from connection
func OpenChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %s", err)
	}

	return ch
}

// Set up queue for messages with it's name and params
// like "durable" for saving queue's state and messages if server restart
func DeclareQueue(name string, ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	return q
}

// Connect consumer, create channel, declare queue and get storage
func ConnectRMQ(cfg *config.Config) *RMQConnection {
	c := &RMQConnection{}
	c.Connection = ConnectAMQP(cfg.Url)
	c.Channel = OpenChannel(c.Connection)
	c.Queue = DeclareQueue("img", c.Channel)

	return c
}

// Take file with id and header request, publish it to queue
func (p *RMQConnection) Publish(id string, file multipart.File, header *multipart.FileHeader) error {
	f, err := ioutil.ReadAll(file)
	log.Println("File readed")
	if err != nil {
		return err
	}

	err = p.Channel.Publish(
		"",           // exchange
		p.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			MessageId:   id,
			ContentType: header.Header.Get("Content-Type"),
			Body:        f,
		},
	)
	if err != nil {
		return err
	}

	log.Println("File published")

	return nil
}

// Start consuming messages and handle it
func (c *RMQConnection) Consume(handler func(string, []byte) error) {
	// Declare Qos for message buffer
	err := c.Channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Fatalf("Failed to set QoS: %s", err)
	}

	// Consume messages from queue
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

	// Handle messages with passed func
	go func() {
		for msg := range msgs {
			log.Println("Consuming")
			time.Sleep(time.Second * 2)
			err = handler(msg.MessageId, msg.Body)
			if err != nil {
				log.Println(err)
			}
		}
	}()
	<-forever
}

func (c *RMQConnection) Disconnect() {
	c.Connection.Close()
	c.Channel.Close()
}
