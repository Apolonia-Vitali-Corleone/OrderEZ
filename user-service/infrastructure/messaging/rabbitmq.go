package messaging

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

const defaultRabbitMQURL = "amqp://guest:guest@192.168.233.136:5672/"

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
func dialRabbit(rawURL string) (*amqp.Connection, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("invalid RABBITMQ_URL %q: %w", rawURL, err)
	}

	q := u.Query()
	if q.Get("heartbeat") == "" {
		q.Set("heartbeat", "10")
	}
	if q.Get("connection_timeout") == "" {
		q.Set("connection_timeout", "10s")
	}
	u.RawQuery = q.Encode()
	switch strings.ToLower(u.Scheme) {
	case "amqp":
		return amqp.Dial(u.String())
	case "amqps":
		{

			host, _, splitErr := net.SplitHostPort(u.Host)
			if splitErr != nil {
				host = u.Host
			}
			tlsCfg := &tls.Config{
				ServerName: host,
			}
			return amqp.DialTLS(u.String(), tlsCfg)
		}
	default:
		return nil, fmt.Errorf("unsupported URL scheme: %s (use amqps:// or amqp://)", u.Scheme)
	}

}

func NewRabbitMQ() (*RabbitMQ, error) {
	raw := rabbitMQURL()
	type dialResult struct {
		conn *amqp.Connection
		err  error
	}
	ch := make(chan dialResult, 1)
	go func() {
		conn, err := dialRabbit(raw)
		ch <- dialResult{conn: conn, err: err}
	}()

	select {
	case res := <-ch:
		if res.err != nil {
			return nil, fmt.Errorf("failed to connect to RabbitMQ %q: %w", raw, res.err)
		}
		conn := res.conn
		ch, err := conn.Channel()
		if err != nil {
			_ = conn.Close()
			return nil, fmt.Errorf("failed to open channel %w", err)
		}
		return &RabbitMQ{conn: conn, channel: ch}, nil

	case <-time.After(30 * time.Second):
		return nil, errors.New("dial RabbitMQ timeout (30s)")

	}
}

func (mq *RabbitMQ) Publish(queue string, body []byte, autoDeclare bool) error {
	if mq == nil || mq.channel == nil {
		return errors.New("rabbitmq: channel is not available")
	}
	if queue == "" {
		return errors.New("rabbitmq: queue name is required")
	}

	if autoDeclare {
		if _, err := mq.channel.QueueDeclare(
			queue,
			true,
			false,
			false,
			false,
			nil,
		); err != nil {
			return fmt.Errorf("declare queue %q: %w", queue, err)
		}
	}

	pub := amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
		Timestamp:   time.Now(),
	}

	if err := mq.channel.Publish(
		"",
		queue,
		false,
		false,
		pub,
	); err != nil {
		return fmt.Errorf("publish to %q: %w", queue, err)
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
