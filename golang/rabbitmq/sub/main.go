package main

import "rabbitmq-demo/rabbitmq"

func main() {
	mq := rabbitmq.NewRabbitMQPubSub("testexchange")
	mq.ConsumerSub()
}
