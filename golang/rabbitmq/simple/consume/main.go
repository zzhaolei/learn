package main

import (
	"fmt"

	"rabbitmq-demo/rabbitmq"
)

func main() {
	mq := rabbitmq.NewRabbitMQSimple("simple")
	defer mq.Destructor()
	mq.ConsumeSimple()
	fmt.Println("done")
}
