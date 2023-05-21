package rabbitmq

import (
	"context"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQRouting(exchange string, key string) *RabbitMQ {
	rabbitmq, err := NewRabbitMQ("", exchange, key)
	failOnError(err, "创建路由模式的 rabbitmq 实例失败")
	return rabbitmq
}

func (mq *RabbitMQ) PublishRouting(message []byte) {
	var err error
	err = mq.exchangeDeclare("direct")
	failOnError(err, "声明 exchange 失败")

	// 发送消息
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

func (mq *RabbitMQ) ConsumerRouting() {
	var err error
	err = mq.exchangeDeclare("direct")
	failOnError(err, "声明 exchange 失败")

	queue, err := mq.channel.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "声明队列失败")

	// 绑定队列
	err = mq.channel.QueueBind(
		queue.Name,
		mq.Key,
		mq.Exchange,
		false,
		nil,
	)
	failOnError(err, " 绑定队列失败")

	// 接收消息
	msgs, err := mq.channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "消费消息失败")

	ch := make(chan struct{})
	go func() {
		for msg := range msgs {
			log.Printf("Recv msg: %s", msg.Body)
		}
	}()

	<-ch
}
