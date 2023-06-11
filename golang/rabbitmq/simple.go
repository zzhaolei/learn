package rabbitmq

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// NewRabbitMQSimple 创建简单模式下 rabbitmq
func NewRabbitMQSimple(queue string) *RabbitMQ {
	rabbitmq, err := NewRabbitMQ(queue, "", "")
	failOnError(err, "创建简单模式的 rabbitmq 实例异常")
	return rabbitmq
}

func (mq *RabbitMQ) PublishSimple(msg []byte) {
	var err error
	// 声明队列
	err = mq.queueDeclare()
	failOnError(err, "声明队列失败")

	// 发送消息
	err = mq.channel.PublishWithContext(
		context.Background(),
		"",
		mq.Queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
	failOnError(err, "发送消息失败")
}

func (mq *RabbitMQ) ConsumeSimple() {
	var err error
	// 声明队列
	err = mq.queueDeclare()
	failOnError(err, "声明队列失败")

	// 消费消息
	msgs, err := mq.channel.Consume(
		mq.Queue,
		"",
		true,
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

	log.Printf("Wait...")
	<-ch
}
