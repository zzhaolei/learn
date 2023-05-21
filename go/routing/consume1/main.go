package main

import "learn/go/rabbitmq"

func main() {
	mq := rabbitmq.NewRabbitMQRouting("texchange", "texchange_two")
	mq.ConsumerRouting()
}
