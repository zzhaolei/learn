package rabbitmq

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQTopic(exchange, key string) *RabbitMQ {
	rabbitmq, err := NewRabbitMQ("", exchange, key)
	failOnError(err, "创建 topic 模式的 rabbitmq 实例失败")
	return rabbitmq
}

func (mq *RabbitMQ) PublishTopic(message []byte) {
	var err error
	err = mq.exchangeDeclare("topic")
	failOnError(err, "声明 exchange 失败")

	err = mq.channel.PublishWithContext(
		context.Background(),
		mq.Exchange,
		mq.Key,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	failOnError(err, "发送消息失败")
}

func (mq *RabbitMQ) ConsumerTopic() {

}
