package messaging

import (
	"errors"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

const defaultRabbitMQURL = "amqp://guest:guest@127.0.0.1:5672/"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func rabbitMQURL() string {
	if url := os.Getenv("RABBITMQ_URL"); url != "" {
		return url
	}
	return defaultRabbitMQURL
}

func NewRabbitMQ() (*RabbitMQ, error) {
	conn, err := amqp.Dial(rabbitMQURL())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}
	return &RabbitMQ{conn: conn, channel: ch}, nil
}

func (mq *RabbitMQ) Publish(queue string, body []byte) error {
	if mq == nil || mq.channel == nil {
		return errors.New("rabbitmq: channel is not available")
	}
	if queue == "" {
		return errors.New("rabbitmq: queue name is required")
	}

	q, err := mq.channel.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue %q: %w", queue, err)
	}

	if err := mq.channel.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	); err != nil {
		return fmt.Errorf("failed to publish message to queue %q: %w", queue, err)
	}
	return nil
}

// Close 关闭 RabbitMQ 连接和通道
func (mq *RabbitMQ) Close() error {
	if mq == nil {
		return nil
	}

	var err error
	if mq.channel != nil {
		if closeErr := mq.channel.Close(); closeErr != nil && !errors.Is(closeErr, amqp.ErrClosed) {
			err = errors.Join(err, closeErr)
		}
	}
	if mq.conn != nil {
		if closeErr := mq.conn.Close(); closeErr != nil && !errors.Is(closeErr, amqp.ErrClosed) {
			err = errors.Join(err, closeErr)
		}
	}
	return err
}
