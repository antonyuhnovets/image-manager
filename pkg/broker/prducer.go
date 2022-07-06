package broker

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/antonyuhnovets/image-manager/pkg/config"
	"github.com/antonyuhnovets/image-manager/pkg/helpers"
)

type Producer struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

func (p *Producer) Connect(cfg *config.Config) {
	p.Connection = helpers.ConnectAMQP(cfg.AMQP_URL)
	p.Channel = helpers.OpenChannel(p.Connection)
	p.Queue = helpers.DeclareQueue("img", p.Channel)
	log.Printf("Producer connected to %s", cfg.AMQP_URL)
}

func (p *Producer) Publish(file multipart.File, header *multipart.FileHeader) {
	f, err := ioutil.ReadAll(file)
	fmt.Println("File readed")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	err = p.Channel.Publish(
		"",           // exchange
		p.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Body: f,
		},
	)
	fmt.Println("File published")
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}
}

func (p *Producer) Disconnect() {
	p.Channel.Close()
	p.Connection.Close()
}
