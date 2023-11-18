package main

import "rabbitmq-demo/rabbitmq"

func main() {
	mq := rabbitmq.NewRabbitMQRouting("texchange", "texchange_one")
	mq.ConsumerRouting()
}
