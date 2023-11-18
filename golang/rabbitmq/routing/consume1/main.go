package main

import "rabbitmq-demo/rabbitmq"

func main() {
	mq := rabbitmq.NewRabbitMQRouting("texchange", "texchange_two")
	mq.ConsumerRouting()
}
