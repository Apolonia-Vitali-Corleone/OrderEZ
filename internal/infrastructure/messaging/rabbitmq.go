package messaging

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ() (*RabbitMQ, error) {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.32.137:5672/")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}
	return &RabbitMQ{conn: conn, channel: ch}, nil
}

func (mq *RabbitMQ) Publish(queue string, body []byte) error {
	q, err := mq.channel.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}

	err = mq.channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}
	log.Printf("[*] Sent %s", body)
	return nil
}

// Close 关闭 RabbitMQ 连接和通道
func (mq *RabbitMQ) Close() {
	mq.channel.Close()
	mq.conn.Close()
}
