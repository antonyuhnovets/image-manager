package helpers

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Connect AMQP server by connecting string
func ConnectAMQP(url string) *amqp.Connection {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect RabbitMQ: %s", err)
	}
	return conn
}

// Open channel by connection
func OpenChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %s", err)
	}
	return ch
}

// Declare queue with passed name
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
