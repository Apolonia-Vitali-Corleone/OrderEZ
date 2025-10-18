// 文件：messaging/rabbitmq.go
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

const (
	// 默认值：本地明文（仅本机测试用）
	defaultRabbitMQURL = "amqp://guest:guest@127.0.0.1:5672/"
)

// RabbitMQ 封装连接与通道
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// 读取连接 URL；优先 RABBITMQ_URL，缺省用默认
func rabbitMQURL() string {
	if u := os.Getenv("RABBITMQ_URL"); u != "" {
		return u
	}
	return defaultRabbitMQURL
}

// 解析并拨号（支持 amqps + SNI）
func dialRabbit(rawURL string) (*amqp.Connection, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("invalid RABBITMQ_URL %q: %w", rawURL, err)
	}

	// 附带心跳 / 连接超时（可通过 URL 查询参数覆盖）
	// 例：?heartbeat=10&connection_timeout=10s
	q := u.Query()
	if q.Get("heartbeat") == "" {
		q.Set("heartbeat", "10") // 秒
	}
	if q.Get("connection_timeout") == "" {
		q.Set("connection_timeout", "10s")
	}
	u.RawQuery = q.Encode()

	switch strings.ToLower(u.Scheme) {
	case "amqps":
		// TLS + SNI
		host, _, splitErr := net.SplitHostPort(u.Host)
		if splitErr != nil {
			host = u.Host
		}
		tlsCfg := &tls.Config{
			ServerName: host, // **关键：SNI 与主机名一致**
			// RootCAs: 使用系统 CA（Amazon MQ 的证书一般是公共 CA）
			// 如用自签证书，可在此加载 RootCAs 或临时 InsecureSkipVerify=true（仅排障）
		}
		return amqp.DialTLS(u.String(), tlsCfg)

	case "amqp":
		return amqp.Dial(u.String())

	default:
		return nil, fmt.Errorf("unsupported URL scheme: %s (use amqps:// or amqp://)", u.Scheme)
	}
}

// NewRabbitMQ 建立连接并打开一个通道
func NewRabbitMQ() (*RabbitMQ, error) {
	raw := rabbitMQURL()

	// 可选：增加一次性拨号超时保护（避免无限卡住）
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
			return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", res.err)
		}
		conn := res.conn
		ch, err := conn.Channel()
		if err != nil {
			_ = conn.Close()
			return nil, fmt.Errorf("failed to open channel: %w", err)
		}
		return &RabbitMQ{conn: conn, channel: ch}, nil

	case <-time.After(30 * time.Second):
		return nil, errors.New("dial RabbitMQ timeout (30s)")
	}
}

// Publish 发布一条消息；可选自动声明队列（durable=true）
func (mq *RabbitMQ) Publish(queue string, body []byte, autoDeclare bool) error {
	if mq == nil || mq.channel == nil {
		return errors.New("rabbitmq: channel is not available")
	}
	if queue == "" {
		return errors.New("rabbitmq: queue name is required")
	}

	if autoDeclare {
		if _, err := mq.channel.QueueDeclare(
			queue, // name
			true,  // durable
			false, // autoDelete
			false, // exclusive
			false, // noWait
			nil,   // args
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
		"",    // default exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		pub,
	); err != nil {
		return fmt.Errorf("publish to %q: %w", queue, err)
	}
	return nil
}

// Close 关闭通道与连接
func (mq *RabbitMQ) Close() error {
	if mq == nil {
		return nil
	}
	var err error
	if mq.channel != nil {
		if e := mq.channel.Close(); e != nil && !errors.Is(e, amqp.ErrClosed) {
			err = errors.Join(err, e)
		}
	}
	if mq.conn != nil {
		if e := mq.conn.Close(); e != nil && !errors.Is(e, amqp.ErrClosed) {
			err = errors.Join(err, e)
		}
	}
	return err
}
