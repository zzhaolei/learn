package main

import "golang/rabbitmq"

func main() {
	mq := rabbitmq.NewRabbitMQRouting("texchange", "texchange_two")
	mq.ConsumerRouting()
}
