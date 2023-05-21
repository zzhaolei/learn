package rabbitmq

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQPubSub(exchange string) *RabbitMQ {
	rabbitmq, err := NewRabbitMQ("", exchange, "")
	failOnError(err, "创建订阅模式的 rabbitmq 实例失败")
	return rabbitmq
}

func (mq *RabbitMQ) PublishPub(message []byte) {
	var err error
	err = mq.exchangeDeclare("fanout")
	failOnError(err, "声明 Exchange 失败")

	// 发送消息
	err = mq.channel.PublishWithContext(
		context.Background(),
		mq.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	failOnError(err, "发送消息到 Exchange 中失败")
}

func (mq *RabbitMQ) ConsumerSub() {
	var err error
	err = mq.exchangeDeclare("fanout")
	failOnError(err, "声明 Exchange 失败 s")

	// 创建队列
	queue, err := mq.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "声明 Queue 失败")

	fmt.Println("随机生成的Queue 名称为", queue.Name)

	err = mq.channel.QueueBind(
		queue.Name,
		"",
		mq.Exchange,
		false,
		nil,
	)
	failOnError(err, "绑定 Queue 失败")

	// 接收消息
	msgs, err := mq.channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "接收消息失败")

	ch := make(chan struct{})
	go func() {
		for msg := range msgs {
			log.Printf("Recv msg: %s", msg.Body)
		}
	}()
	log.Println("Wait...")
	<-ch
}
