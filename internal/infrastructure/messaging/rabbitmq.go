package messaging

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	conn *amqp.Connection
}

func NewRabbitMQ(conn *amqp.Connection) *RabbitMQ {
	return &RabbitMQ{conn: conn}
}

func InitRabbitMQ() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	return conn
}

func (r *RabbitMQ) Publish(queueName string, body []byte) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	return err
}
