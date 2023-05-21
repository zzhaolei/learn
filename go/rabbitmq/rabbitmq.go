package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// amqp://用户:密码@服务地址:端口/virtual-host
const mq_url = "amqp://rabbit:rabbit@127.0.0.1:5672/rabbithost"

type RabbitMQ struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	Queue    string
	Exchange string
	Key      string
}

// NewRabbitMQ 创建一个 mq 对象，有可能创建失败 s
func NewRabbitMQ(queue, exchange, key string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(mq_url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn:     conn,
		channel:  channel,
		Queue:    queue,
		Exchange: exchange,
		Key:      key,
	}, nil
}

// Destructor 关闭存在的 conn 和 channel 连接 s
func (mq *RabbitMQ) Destructor() {
	// NewRabbitMQ 保证 conn 和 channel 存在
	mq.conn.Close()
	mq.channel.Close()
}

// queueDeclare 声明队列，如果队列存在则跳过
func (mq *RabbitMQ) queueDeclare() (err error) {
	_, err = mq.channel.QueueDeclare(
		mq.Queue,
		false,
		false,
		false,
		false,
		nil,
	)
	return
}

func (mq *RabbitMQ) exchangeDeclare(kind string) (err error) {
	return mq.channel.ExchangeDeclare(
		mq.Exchange,
		kind,
		true,
		false,
		false,
		false,
		nil,
	)
}
