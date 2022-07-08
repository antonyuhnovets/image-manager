package broker

import (
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

// Connect producer, create channel and declare queue
func (p *Producer) Connect(cfg *config.Config) {
	p.Connection = helpers.ConnectAMQP(cfg.AMQP_URL)
	p.Channel = helpers.OpenChannel(p.Connection)
	p.Queue = helpers.DeclareQueue("img", p.Channel)
	log.Println("Producer connected")
}

// Publish message with image
func (p *Producer) Publish(id string, file multipart.File, header *multipart.FileHeader) error {
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
	log.Println("File published")
	if err != nil {
		return err
	}

	return nil
}

// Disconnect producer
func (p *Producer) Disconnect() {
	p.Channel.Close()
	p.Connection.Close()
}