package main

import "golang/rabbitmq"

func main() {
	mq := rabbitmq.NewRabbitMQPubSub("testexchange")
	mq.ConsumerSub()
}
