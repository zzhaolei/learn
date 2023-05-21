package main

import "learn/go/rabbitmq"

func main() {
	mq := rabbitmq.NewRabbitMQPubSub("testexchange")
	mq.ConsumerSub()
}
